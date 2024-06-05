import { useContext } from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";

import HomePage from "../../Pages/HomePage";
import LogPage from "../../Pages/LogPage";
import Navbar from "../Navbar/Navbar";
import { UserContext } from "../AppContext";

const Routers = () => {
    const { userid } = useContext(UserContext);

    return (
        <div>
            <Router>
                {userid ? <Navbar /> : null}
                <div className="main-content">
                    <Routes>
                        <Route path="/" element={userid ? <HomePage/> : <Navigate to="/log"/>} />
                        <Route path="/log" element={userid ? <Navigate to="/"/> : <LogPage/>} />
                        <Route path="*" element={<Navigate to="/" />} />
                    </Routes>
                </div>
            </Router>
        </div>
    );
};

export default Routers;