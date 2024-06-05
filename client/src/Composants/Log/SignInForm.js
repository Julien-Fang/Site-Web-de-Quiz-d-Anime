import { useEffect, useState } from "react";
import axios from "axios";
import { useContext } from "react";
import { UserContext } from "../AppContext";
import styled from "styled-components";

const Logo = styled.img`
    border-radius: 50%;
    height: 100px;
    width: 100px;
`;

const SignInForm = () => {
    const [id, setId] = useState("");
    const [mdp, setMdp] = useState("");
    const [isReset, setIsReset] = useState(false);

    const loginError = document.getElementById("login-error");
    const mdpError = document.getElementById("mdp-error");
    const headerError = document.getElementById("header-error");
    const myForm = document.getElementById("myForm");


    const { handleLogin } = useContext(UserContext);

    const submitForm = async (e) => {
        e.preventDefault();

        loginError.innerHTML = " ";
        mdpError.innerHTML = " ";
        headerError.innerHTML = " ";
        resetForm();

        await axios({
            method: "POST",
            url: "http://localhost:8000/user/login",
            withCredentials: true,
            data: {
                login: id, 
                password: mdp
            }
        })
        .then(res => {
            handleLogin(res.data.userid);
        })
        .catch(err => {
            if (err.response.data === "Identifiant erroné\n") {
                loginError.innerHTML = err.response.data;
            }
            else if (err.response.data === "Mot de passe erroné\n") {
                mdpError.innerHTML = err.response.data;
            }
            else if (err.response.data === "Champs manquants\n"){
                headerError.innerHTML = err.response.data;
            }
            else {
                console.log(err);
            }
        })
    }

    function resetForm() {
        myForm.querySelectorAll('input').forEach((el) => {
          el.value = '';
        });
        setIsReset(true);
    }


    useEffect(() => {
        if (isReset) {
            setId("");
            setMdp("");
            setIsReset(false);
        }
    }, [isReset]);


    return (
        <div className="login d-flex" style={{flexDirection: "column"}}>
            <p id="header-error" style={{color: "red"}} className="form-text"></p>
            {/* <div className="d-flex justify-content-center">
                <Logo src="/uploads/default.jpg" alt="logo" className="mb-3"/>
            </div>
            <h2 style={{fontSize: "2.5rem"}}>Welcome to Prism</h2> */}
            
            <form method="POST" action="" onSubmit={submitForm} id="myForm">
                <div className="mb-3">
                    <label htmlFor="username" className="form-label text-light">Login</label>
                    <input type="text" className="form-control" name="username" id="username" value={id} onChange={e => {setId(e.target.value)}}/>
                </div>
                <p className="form-text" style={{color: "red"}} id="login-error"></p>
                <div className="mb-3">
                    <label htmlFor="password" className="form-label text-light">Password</label>
                    <input type="password" className="form-control" name="password" id="password" value={mdp} onChange={e => {setMdp(e.target.value)}}/>
                </div>
                <p className="form-text" style={{color: "red"}} id="mdp-error"></p>
                <div className="d-flex justify-content-end">
                    <button type="submit" style={{borderRadius: "10px", height: "40px", width: "140px", fontSize: "1.1rem"}} className="btn btn-primary">Submit</button>
                </div>
            </form>
        </div>
    );

};

export default SignInForm;