FROM {{ .Base }}
RUN the date today is: {{ now | htmlDate }}