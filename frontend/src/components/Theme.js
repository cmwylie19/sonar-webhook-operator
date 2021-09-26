import React from "react";
import { createMuiTheme, ThemeProvider } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import { pink, lightGreen, grey } from "@material-ui/core/colors";

export const pinkTheme = createMuiTheme({
  typography: {
    fontFamily: "Red Hat Text, sans-serif",
  },
  palette: {
    mode: "dark",
    primary: { main: grey[900], contrastText: grey[50] },
    secondary: { main: grey[900] },
    type: "dark",
  },
});

export const greenTheme = createMuiTheme({
  typography: {
    fontFamily: "Red Hat Text, sans-serif",
  },
  palette: {
    primary: { main: grey[900], contrastText: "#fbfbfb" },
    secondary: { main: lightGreen[500] },
    type: "light",
  },
  props: {
    MuiButtonBase: {
      disableRipple: true,
    },
  },

  MuiButton: {
    text: {
      color: lightGreen,
    },
  },
  MuiCssBaseline: {
    "@global": {
      "@font-face": `'Red Hat Display', sans-serif;`,
    },
  },
  MuiStepIcon: {
    root: {
      "&$active": {
        fill: lightGreen,
        "& $text": {
          fill: lightGreen,
        },
      },
    },
    text: {
      fill: lightGreen,
    },
  },
});
export default function ThemeContainer({ children }) {
  return (
    <ThemeProvider theme={pinkTheme}>
      <CssBaseline />
      {children}
    </ThemeProvider>
  );
}
