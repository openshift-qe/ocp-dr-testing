FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.17-openshift-4.10
RUN mkdir -p /go/src/github.com/openshift-qe/ocp-dr-testing
WORKDIR /go/src/github.com/openshift-qe/ocp-dr-testing
COPY . .
