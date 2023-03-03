{{ define "LAYOUT" }}

<html>
<body>
{{ template "CONTENT" . }}
</body>
</html>
{{ end }}

{{ define "CONTENT" }}
Unsafe content: {{ .Attended }}
{{- range .lines }}
{{ . }}
{{- end }}
{{ end }}

{{ template "LAYOUT" . }}
{{ .Animal }}

<!-- {{- range .lines}}
{{.}}
{{end}} -->
