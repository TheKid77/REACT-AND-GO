import HomeIcon from "@mui/icons-material/Home";
import LogoutIcon from "@mui/icons-material/Logout";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import SettingsIcon from "@mui/icons-material/Settings";
import {
  IconButton,
  ListItemIcon,
  Menu,
  MenuItem,
  Typography,
} from "@mui/material";
import React, { useState } from "react";

const NiceMenu = () => {
  const [anchorEl, setAnchorEl] = useState(null);
  const open = Boolean(anchorEl);

  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <IconButton
        onClick={handleClick}
        size="large"
        aria-controls={open ? "menu" : undefined}
        aria-haspopup="true"
        aria-expanded={open ? "true" : undefined}
        style={{ color: "#1976d2" }} // Customize button color
      >
        <MoreVertIcon fontSize="large" />
      </IconButton>

      <Menu
        id="menu"
        anchorEl={anchorEl}
        open={open}
        onClose={handleClose}
        PaperProps={{
          style: {
            borderRadius: "10px",
            minWidth: "200px",
            boxShadow: "0px 5px 10px rgba(0,0,0,0.1)",
          },
        }}
        MenuListProps={{
          style: {
            padding: 0,
          },
        }}
      >
        <MenuItem onClick={handleClose}>
          <ListItemIcon>
            <HomeIcon fontSize="small" />
          </ListItemIcon>
          <Typography variant="inherit">Home</Typography>
        </MenuItem>
        <MenuItem onClick={handleClose}>
          <ListItemIcon>
            <SettingsIcon fontSize="small" />
          </ListItemIcon>
          <Typography variant="inherit">Settings</Typography>
        </MenuItem>
        <MenuItem onClick={handleClose}>
          <ListItemIcon>
            <LogoutIcon fontSize="small" />
          </ListItemIcon>
          <Typography variant="inherit">Logout</Typography>
        </MenuItem>
      </Menu>
    </div>
  );
};

export default NiceMenu;
