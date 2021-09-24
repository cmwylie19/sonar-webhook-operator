import React, {useState} from "react";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import Checkbox from "@mui/material/Checkbox";
import IconButton from "@mui/material/IconButton";
import CommentIcon from "@mui/icons-material/Comment";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { green } from "@material-ui/core/colors";
import Title from "./Title";

const theme = createTheme({
  typography: {
    fontFamily: "Red Hat Display",
  },
  palette: {
      type:"dark",
    secondary: { main: green[400] },
  },
});

export default function ResultList() {
  const [checked, setChecked] = useState([0]);
  const [clicked, setClicked] = useState()
  const handleToggle = (value) => () => {
    const currentIndex = checked.indexOf(value);
    const newChecked = [...checked];

    if (currentIndex === -1) {
      newChecked.push(value);
    } else {
      newChecked.splice(currentIndex, 1);
    }

    setChecked(newChecked);
  };

  return (
    <ThemeProvider theme={theme} >
        <Title>All Analyses</Title>
      <List sx={{ width: "100%", maxWidth: "100%" }} id="resultlist">
        {[0, 1, 2, 3,4,5,6,7,8,9].map((value) => {
          const labelId = `checkbox-list-label-${value}`;

          return (
            <ListItem
            style={{backgroundColor:clicked===value?"#fbfbfb":"inherit"}}
              key={value}
              secondaryAction={
                <IconButton edge="end" aria-label="comments" onClick={()=>setClicked(value)}>
                  <CommentIcon />
                </IconButton>
              }
              disablePadding
            >
              <ListItemButton
                role={undefined}
                onClick={handleToggle(value)}
                dense
              >
                <ListItemIcon>
                  <Checkbox
                    color="secondary"
                    edge="start"
                    checked={checked.indexOf(value) !== -1}
                    tabIndex={-1}
                    disableRipple
                    inputProps={{ "aria-labelledby": labelId }}
                  />
                </ListItemIcon>
                <ListItemText id={labelId} primary={`Line item ${value + 1}`} />
              </ListItemButton>
            </ListItem>
          );
        })}
      </List>
    </ThemeProvider>
  );
}
