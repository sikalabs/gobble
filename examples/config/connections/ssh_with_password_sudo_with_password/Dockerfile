FROM sikalabs/sshd
RUN \
  apt-get update && \
  apt-get install -y --no-install-recommends \
    sudo && \
  rm -rf /var/lib/apt/lists
RUN mkdir -p /home/demo
RUN useradd demo -s /bin/bash
RUN chown demo:demo /home/demo
RUN echo "demo:demo" | chpasswd
RUN echo "demo ALL=(ALL) ALL" >> /etc/sudoers
