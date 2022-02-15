import React from 'react'
import { BrowserRouter as Router, Navigate  ,Route, Routes } from "react-router-dom";
import Calculadora from "./Componentes/Calculadora.js"

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/Calculadora" element={<Calculadora />} />
        <Route path="/" element={<Navigate replace to="/Calculadora" />} />
      </Routes>
    </Router>
  );
}

export default App;
