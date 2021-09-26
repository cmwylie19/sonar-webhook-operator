import * as React from "react";
import PropTypes from "prop-types";
import Typography from "@mui/material/Typography";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { green } from "@material-ui/core/colors";

const theme = createTheme({
  typography: {
    fontFamily: "Red Hat Display",
  },
  palette: {
    type: "dark",
    secondary: { main: green[400] },
  },
});
function Title(props) {
  return (
    <ThemeProvider theme={theme}>
      <Typography component="h2" variant="h6" color="secondary" gutterBottom>
        {props.children}
      </Typography>
    </ThemeProvider>
  );
}

Title.propTypes = {
  children: PropTypes.node,
};

export default Title;
