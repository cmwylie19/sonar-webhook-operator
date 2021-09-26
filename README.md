# Sonar Webhook Operator
<!-- ![Logo](logo-512.png) -->
<img src="logo-512.png" alt="drawing" width="150"/>

Project Components
- [Webhook Backend](#webhook)
- [Operator](#operator)
- [GitHub Action](#github-action)

## Webhook
_The webhook will be triggered from a HTTP POST request from sonarqube._

## Operator
_The operator will control the Webhook's `Deployment`, `Service`, and `ServiceAccount`._

## GitHub Action
_The GitHub Action will query the webhook to get the sonarqube results. The action will pass/fail based on quality gate results for a given job._

