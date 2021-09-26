import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { grey } from "@material-ui/core/colors";

const theme = createTheme({
  palette: {
    primary: { main: grey[900], contrastText: grey[50] },
    secondary: { main: grey[900] },
    type: "dark",
  },
});
export default function Nav() {
  return (
    <ThemeProvider theme={theme}>
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="static">
          <Toolbar>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              Sonar Webhook
            </Typography>

            <Typography variant="h6" component="div">
              Docs
            </Typography>
          </Toolbar>
        </AppBar>
      </Box>
    </ThemeProvider>
  );
}
