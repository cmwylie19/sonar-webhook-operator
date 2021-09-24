const scanner = require("sonarqube-scanner");
scanner(
  {
    serverUrl: "http://34.75.152.50:9000/",
    options: {
      "sonar.sources": "src",
      "sonar.analysis.buildNumber":"69",
      "sonar.exclusions": "**/__tests__/**",
      "sonar.tests": "src/__tests__",
      "sonar.javascript.lcov.reportPaths": "coverage/lcov.info",
      "sonar.projectKey": "sonar-webhook-frontend",
      "sonar.projectName": "sonar-webhook-frontend",
      "sonar.sourceEncoding": "UTF-8",
      "sonar.projectVersion": "1.0",
      "sonar.javascript.file.suffixes": ".js",
      "sonar.login": "ef6df74928660fb2d034eeda6c0dccf3638e6320"
    },
  },
  (code) => process.exit(code)
);
