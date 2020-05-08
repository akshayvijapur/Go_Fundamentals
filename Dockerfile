FROM ubuntu:18.04
RUN apt-get update \
    && apt-get install -y python3-pip python3-dev git wget \
    && cd /usr/local/bin \
    && ln -s /usr/bin/python3 python \
    && pip3 install --upgrade pip
WORKDIR /root/
RUN git clone https://github.com/keithito/tacotron.git
WORKDIR /root/tacotron
RUN wget https://data.keithito.com/data/speech/LJSpeech-1.1.tar.bz2
RUN tar -xvf LJSpeech-1.1.tar.bz2
RUN pip3 install intel-tensorflow==1.15.2
RUN pip3 install scipy==1.1.0
RUN pip3 install scikit-learn==0.21.3
RUN pip3 install falcon==1.2.0
RUN pip3 install inflect==0.2.5
RUN pip3 install librosa==0.5.1
RUN pip3 install tqdm
RUN pip3 install scipy --upgrade
RUN pip3 install numba==0.48.0
RUN pip3 install setuptools>=41.0.0
RUN pip3 install tensorboard==1.15.0
RUN python3 preprocess.py --dataset ljspeech
RUN python3 train.py
