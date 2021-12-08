
import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import "./index.css";

import Car from "./components/Car";
import List from "./components/List";
import UI from "./components/UL";
import Datatable from "./components/Table";

ReactDOM.render(
  <>
    <Datatable />
    <Car age="25" eye="balck" />
    <List />
    <UI />
  </>, 
  document.getElementById('root')
);