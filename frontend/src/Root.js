import React, {useEffect} from "react";
import Nav from "./components/Nav";
import Theme from "./components/Theme";
import ResultList from "./components/ResultList";
import ResultTabs from "./components/ResultTabs"
import Container from "@mui/material/Container";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";

function App() {
  return (
    <Theme>
      <Nav />
      <Container component="main" maxWidth="lg" sx={{ mt: 4, mb: 4,marginTop:"50px" }}>
        <Grid container spacing={7}>
          <Grid item xs={12} md={8} lg={6}>
            {/* <Paper
            elevation={12}
            color="primary"
              sx={{
                p: 2,
                display: "flex",
                flexDirection: "column",
                width: "400px",
                backgroundColor: "#676767"
              }}
            > */}
              <ResultList />
            {/* </Paper> */}
          </Grid>
          <Grid item xs={12} md={12} lg={6}>
            <Paper
            elevation={9}
            color="primary"
              sx={{
                p: 2,
                display: "flex",
                flexDirection: "column",
                width: "100%"
              }}
            >
              <ResultTabs />
            </Paper>
          </Grid>
        </Grid>
      </Container>
    </Theme>
  );
}

export default App;
