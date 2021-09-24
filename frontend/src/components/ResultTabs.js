import * as React from 'react';
import PropTypes from 'prop-types';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { makeStyles } from '@mui/styles';

function TabPanel(props) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          <Typography>{children}</Typography>
        </Box>
      )}
    </div>
  );
}

TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.number.isRequired,
  value: PropTypes.number.isRequired,
};

function a11yProps(index) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

const useStyles = makeStyles({
    root: {
        "& .MuiTabs-indicator": {
            // display: "none"
            backgroundColor: "orange"
          },
          "& .MuiTab-textColorPrimary":{
              color: "green"
          },
          "& .css-1h9z7r5-MuiButtonBase-root-MuiTab-root":{
              color: "red"
          },
          "& .css-1h9z7r5-MuiButtonBase-root-MuiTab-root.Mui-selected": {
              color: "green"
          }
    }, // a style rule
    label: {}, // a nested style rule
    
  });

export default function ResultTabs(props) {
  const [value, setValue] = React.useState(0);
    const classes = useStyles()
  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  return (
    <Box sx={{ width: '100%' }} id="tabs">
      <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
        <Tabs value={value} onChange={handleChange} aria-label="basic tabs example" className={classes.root}>
          <Tab label="Item One" {...a11yProps(0)} />
          <Tab label="Item Two" {...a11yProps(1)} />
          <Tab label="Item Three" {...a11yProps(2)} />
        </Tabs>
      </Box>
      <TabPanel value={value} index={0}>
        Item One
      </TabPanel>
      <TabPanel value={value} index={1}>
        Item Two
      </TabPanel>
      <TabPanel value={value} index={2} id="two">
        Item Three
      </TabPanel>
    </Box>
  );
}