import { Routes, Route } from "react-router-dom";

import "./App.css";
import Home from "./page/Home";
import Blog from "./page/Blog";
import Add from "./page/Add";
import Edit from "./page/Edit";
import Delete from "./page/Delete";
import Header from "./component/layout/Header";
import Footer from "./component/layout/Footer";

function App() {
  return (
    <>
      <Header></Header>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/add" element={<Add />} />
        <Route path="/edit/:id" element={<Edit />} />
        <Route path="/delete/:id" element={<Delete />} />
        <Route path="/blog/:id" element={<Blog />} />
      </Routes>
      <Footer></Footer>
    </>
  );
}

export default App;