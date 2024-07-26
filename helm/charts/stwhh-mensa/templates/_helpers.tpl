{{/*
Expand the name of the chart.
*/}}
{{- define "stwhh-mensa.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "stwhh-mensa.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{- define "stwhh-mensa.frontend.fullname" -}}
{{- printf "%s-frontend" (include "stwhh-mensa.fullname" . ) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "stwhh-mensa.backend.fullname" -}}
{{- printf "%s-backend" (include "stwhh-mensa.fullname" . ) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "stwhh-mensa.frontend.name" -}}
{{- printf "%s-frontend" (include "stwhh-mensa.name" . ) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "stwhh-mensa.backend.name" -}}
{{- printf "%s-backend" (include "stwhh-mensa.name" . ) | trunc 63 | trimSuffix "-" }}
{{- end }}


{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "stwhh-mensa.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "stwhh-mensa.labels" -}}
helm.sh/chart: {{ include "stwhh-mensa.chart" . }}
{{ include "stwhh-mensa.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "stwhh-mensa.frontend.labels" -}}
helm.sh/chart: {{ include "stwhh-mensa.chart" . }}
{{ include "stwhh-mensa.frontend.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "stwhh-mensa.backend.labels" -}}
helm.sh/chart: {{ include "stwhh-mensa.chart" . }}
{{ include "stwhh-mensa.backend.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "stwhh-mensa.selectorLabels" -}}
app.kubernetes.io/name: {{ include "stwhh-mensa.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "stwhh-mensa.frontend.selectorLabels" -}}
{{ include "stwhh-mensa.selectorLabels" . }}
app.kubernetes.io/component: frontend
{{- end }}

{{- define "stwhh-mensa.backend.selectorLabels" -}}
{{ include "stwhh-mensa.selectorLabels" . }}
app.kubernetes.io/component: backend
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "stwhh-mensa.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "stwhh-mensa.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}