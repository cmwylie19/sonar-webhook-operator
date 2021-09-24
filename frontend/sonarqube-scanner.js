const scanner = require("sonarqube-scanner");
scanner(
  {
    serverUrl: "http://localhost:9000",
    options: {
      "sonar.sources": "src",
      "sonar.analysis.buildNumber":"69",
      "sonar.exclusions": "**/__tests__/**",
      "sonar.tests": "src/__tests__",
      "sonar.javascript.lcov.reportPaths": "coverage/lcov.info",
      "sonar.projectKey": "sonar-webhook-frontend",
      "sonar.projectName": "Sonar Webhook Frontend",
      "sonar.sourceEncoding": "UTF-8",
      "sonar.projectVersion": "1.0",
      "sonar.javascript.file.suffixes": ".js",
      "sonar.login": process.env.TOKEN
    },
  },
  (code) => process.exit(code)
);
