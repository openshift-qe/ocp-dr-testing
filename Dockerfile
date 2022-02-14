FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.17-openshift-4.10 AS builder
RUN mkdir -p /go/src/github.com/openshift-qe/ocp-dr-testing
WORKDIR /go/src/github.com/openshift-qe/ocp-dr-testing
COPY . .
FROM registry.ci.openshift.org/ocp/4.10:tools
RUN sh -c 'echo -e "[google-cloud-sdk]\nname=Google Cloud SDK\nbaseurl=https://packages.cloud.google.com/yum/repos/cloud-sdk-el7-x86_64\nenabled=1\ngpgcheck=1\nrepo_gpgcheck=1\ngpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg\n       https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg" > /etc/yum.repos.d/google-cloud-sdk.repo' && \
    PACKAGES="google-cloud-sdk git gzip zip util-linux openssh-clients httpd-tools make gcc" && \
    yum update -y && \
    yum install --setopt=tsflags=nodocs -y $PACKAGES && yum clean all && rm -rf /var/cache/yum/* && \
    git config --system user.name test-private && \
    git config --system user.email test-private@test.com && \
    chmod g+w /etc/passwd && \
    rm -rf /root/.config/gcloud
RUN pip3 install dotmap minio pyyaml requests
RUN curl -s -k https://dl.google.com/go/go1.17.6.linux-amd64.tar.gz -o go1.17.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz && rm -fr go1.17.6.linux-amd64.tar.gz && \
    mkdir -p /tmp/goproject && mkdir -p /tmp/gocache && chmod -R g=u /tmp/goproject && \
    chmod -R g+rw /tmp/goproject && chmod -R g=u /tmp/gocache && chmod -R g+rw /tmp/gocache 
RUN curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/ocp/stable/openshift-client-linux.tar.gz | tar -xvzf -  &&  mv oc /bin && mv kubectl /bin  
