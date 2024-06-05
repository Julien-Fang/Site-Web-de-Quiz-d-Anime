import { useState, useEffect } from 'react';
import axios from 'axios';
import SignInForm from './SignInForm';

const SignUpForm = () => {
    const [id, setId] = useState("")
    const [mdp, setMdp] = useState("")
    const [confirm, setConfirm] = useState("")
    const [pseudo, setPseudo] = useState("")

    const [isReset, setIsReset] = useState(false)
    const [isSubmit, setIsSubmit] = useState(false)

    const headerError = document.getElementById('header-error');
    const myForm = document.getElementById('myForm');

    const submitForm = (e) => {
        e.preventDefault();

        headerError.innerHTML = "";

        if (mdp !== confirm) {
            resetForm();
            headerError.innerHTML = "Les mots de passe ne correspondent pas";
            return;
        }

        axios({
            method: 'post',
            url: 'http://localhost:8000/user/signup',
            withCredentials: true,
            data: {
                user: {
                    login: id,
                    password: mdp,
                },
                pseudo: pseudo,
            }
        })
        .then(res => {
            setIsSubmit(true);
        })
        .catch(err => {
            console.log(err.response.data)
            headerError.innerHTML = err.response.data;
            resetForm();
        });
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
            setConfirm("");
            setPseudo("");
            setIsReset(false);
        }
    }, [isReset]);

    return (
        isSubmit 
            ? 
            (<>
                <h3 className="form-text text-light">User successfully created, Welcome, You can now Log In</h3>
                <SignInForm /> 
            </>)
            :
            (<div className="signin">
                <p id="header-error" style={{color: "red"}} className="form-text"></p>
                <h2 className='text-light' style={{fontSize: "2.2rem"}}>Create Your Account</h2>

                <form method="POST" action="" onSubmit={submitForm} id='myForm'>

                    <div className="mb-3">
                        <label htmlFor="identifiant" className="form-label text-light">Login</label>
                        <input type="text" className="form-control" name="identifiant" id="identifiant" value={id} onChange={e => {setId(e.target.value)}}/>
                    </div>

                    <div className="mb-3">
                        <label htmlFor="password" className="form-label text-light">Password</label>
                        <input type="password" className="form-control" name="password" id="password" value={mdp} onChange={e => {setMdp(e.target.value)}}/>
                    </div>

                    <div className="mb-3">
                        <label htmlFor="confirm" className="form-label text-light">Confirm Password</label>
                        <input type="password" className="form-control" name="confirm" id="confirm" value={confirm} onChange={e => {setConfirm(e.target.value)}}/>
                    </div>

                    <br/>

                    <div className="mb-3">
                        <label htmlFor="pseudo" className='form-label text-light'>Pseudo</label>
                        <input type="text" className='form-control' name="pseudo" id="pseudo" value={pseudo} onChange={e => {setPseudo(e.target.value)}}/>
                    </div>

        
                    <div className="d-flex justify-content-end">
                        <button type="submit" style={{borderRadius: "10px", height: "40px", width: "140px", fontSize: "1.1rem"}} className="btn btn-primary">Register</button>
                    </div>
                </form>
            </div>)
    );
};

export default SignUpForm;