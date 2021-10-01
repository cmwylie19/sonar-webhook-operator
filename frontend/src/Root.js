import React, { useEffect, useState } from "react";
import Nav from "./components/Nav";
import Theme from "./components/Theme";
import ResultList from "./components/ResultList";
import ResultTabs from "./components/ResultTabs";
import FilterField from "./components/FilterField";
import Container from "@mui/material/Container";
import Grid from "@mui/material/Grid";
import API from "./utils";
function App() {
  const [sonarResults, setSonarResults] = useState([]);
  const [token, setToken] = useState();
  useEffect(() => {
    const api = new API();
    return api
      .fetchResults()
      .then((response) => setSonarResults(response.data));
    // Check token
    // if no token, redirect to login
  }, []);

  return (
    <Theme>
      <Nav />
      <Container
        component="main"
        maxWidth="lg"
        sx={{ mt: 4, mb: 4, marginTop: "50px" }}
      >
        <Grid container spacing={7}>
          <Grid item xs={12} large={12}>
            {JSON.stringify(sonarResults, undefined, 2)}
            <FilterField />
          </Grid>
          <Grid item xs={12} md={8} lg={6}>
            <ResultList sonarResults={sonarResults} />
          </Grid>
          <Grid item xs={12} md={12} lg={6}>
            <ResultTabs sonarResults={sonarResults} />
          </Grid>
        </Grid>
      </Container>
    </Theme>
  );
}

export default App;
