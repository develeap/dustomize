FROM {{ .Base }}
RUN apt install {{ .git.packageName }}={{ .git.packageVersion }}
RUN echo {{ randAlphaNum 20 }}