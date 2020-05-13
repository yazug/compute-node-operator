FROM registry.access.redhat.com/ubi8/ubi-minimal:latest as build-env

RUN microdnf -y install golang

WORKDIR /go/src/github.com/openstack-k8s-operators/compute-node-operator
ADD . .

RUN CGO_ENABLED=0 go build -mod=vendor -a -ldflags '-extldflags "-static"' -o /csv-generator tools/csv-generator.go
RUN CGO_ENABLED=0 go build -mod=vendor -a -ldflags '-extldflags "-static"' -o /compute-node-operator cmd/manager/main.go

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

LABEL   name="compute-node-operator" \
        version="1.0" \
        summary="Compute Node Operator" \
        io.k8s.name="compute-node-operator" \
        io.k8s.description="This container includes the compute-node-operator"

ENV OPERATOR=/usr/local/bin/compute-node-operator \
    USER_UID=1001 \
    USER_NAME=compute-node-operator \
    OPERATOR_BUNDLE=/usr/share/compute-node-operator/bundle/

# install operator binary
COPY --from=build-env /compute-node-operator ${OPERATOR}
COPY --from=build-env /csv-generator /usr/local/bin/csv-generator

COPY build/bin /usr/local/bin
COPY bindata /bindata
RUN  /usr/local/bin/user_setup

# install CRDs and required roles, services, etc
RUN  mkdir -p ${OPERATOR_BUNDLE}
COPY deploy/crds/*crd.yaml ${OPERATOR_BUNDLE}

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
