FROM {{ .Base }}
RUN apt install jq
RUN echo {{ randAlphaNum 20 }}