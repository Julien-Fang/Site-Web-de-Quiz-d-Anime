import { useState } from 'react';
import SignUpForm from './SignUpForm';
import SignInForm from './SignInForm';
import { Button } from 'react-bootstrap';

const Log = () => {
    const [mode, setMode] = useState(true);

    const handleMode = (e) => {
        if (e.target.id === 'register') setMode(false);
        else setMode(true);
    };

    return (
        <div className='connection-form' style={{height: "100vh"}}>
            <div className="d-grid col justify-content-center">
                
                <div className="d-flex justify-content-between" style={{marginTop: "10%"}}>
                    <Button variant="primary" 
                        style={{ borderRadius: "10px", height: "40px", width: "140px", fontSize: "1.1rem", marginRight: "10px" }} 
                        onClick={handleMode}
                        id='login'
                    >
                        Log In
                    </Button>
                    <Button variant="primary" 
                        style={{ borderRadius: "10px", height: "40px", width: "140px", fontSize: "1.1rem", marginLeft: "10px"}} 
                        onClick={handleMode}
                        id='register'
                    >
                        Sign In
                    </Button>
                </div>

                {!mode && <div className=''><SignUpForm/></div>}
                {mode && <div className=''><SignInForm /></div>}

            </div>
        </div>
    );
}

export default Log;