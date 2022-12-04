FROM ubuntu:22.10

ARG USER=oasis
ARG UID=1000

LABEL \
  description="Dev workspace with standard tools and packages useful for common work" \
  name="Dev workspace" \
  user-name=$USER \
  user-id=$UID

WORKDIR /workspace

ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Etc/UTC

RUN apt-get update && apt-get install -y \
  locales \
  git \
  sudo \
  unzip \
  vim \
  wget \
  ca-certificates \
  curl \
  gnupg \
  lsb-release \
  tar \
  python3.10 \
  dos2unix \
  iputils-ping;

RUN apt-get install -y ssh

# Add docker GPG key and install docker
RUN  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg; \
  echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null; \
  apt-get update; \
  apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

# Install HELM
RUN curl https://baltocdn.com/helm/signing.asc | sudo apt-key add -; \
  sudo apt-get install apt-transport-https --yes; \
  echo "deb https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list; \
  sudo apt-get update; \
  sudo apt-get install -y helm

# Install kubectl (with binary verification)
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"; \
  curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"; \
  echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check; \
  install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl;


# Install OC
RUN wget --no-check-certificate https://github.com/openshift/origin/releases/download/v3.11.0/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz; \
  tar -xzf openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz; \
  ln -s /workspaces/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit/oc /bin/oc

# Install kubeval
RUN wget https://github.com/instrumenta/kubeval/releases/latest/download/kubeval-linux-amd64.tar.gz; \
  tar xf kubeval-linux-amd64.tar.gz; \
  cp kubeval /usr/local/bin;

# Install minikube
RUN curl -LO "https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64"; \
  install minikube-linux-amd64 /usr/local/bin/minikube;
# Configure minikube
RUN minikube config set driver docker;\
  minikube config set cpus 4;\
  minikube config set memory 8192

# Create user group and add user
RUN groupadd --gid $UID $USER && \
  adduser --shell /bin/bash --uid $UID --gid $UID $USER && \
  echo "$USER ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers; \
  usermod -aG docker $USER

# Install Golang
ARG GO_VERSION=go1.19.3
RUN wget https://go.dev/dl/$GO_VERSION.linux-amd64.tar.gz \
  && tar -C /usr/local -xzf $GO_VERSION.linux-amd64.tar.gz;\
  echo "export PATH=$PATH:/usr/local/go/bin" >> /home/$USER/.bashrc

USER $UID

# setup Nodejs
RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.2/install.sh | bash; \
  export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"; \
  [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm; \
  nvm install 18.12.1;

WORKDIR /workspace