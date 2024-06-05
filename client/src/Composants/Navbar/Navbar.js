import { useContext } from 'react';
import { UserContext } from '../AppContext';
import Logout from './Logout';
import AllPins from './AllPins';


const Navbar = () => {
    const { userid, pseudo, handleLogout } = useContext(UserContext);

    const toHome = () => {
        const toHome = () => {
            try {
                window.location.href = "http://localhost:3000";
            }
            catch (err) {
                handleLogout();
            }
        }
        return (
            <div className="d-flex" onClick={toHome}>
                Home
            </div>
        );
    }

    return (
        userid 
        ? (
            <nav className="navbar nav-bg navbar-expand-lg nav nav-tabs">
                <button className="nav-link col-2" >{toHome()}</button>
                <AllPins/>
                <div className='d-flex align-items-center justify-content-center ms-auto'>
                    <p style={{marginBottom: "0px", marginRight: "2rem"}}>Coucou {pseudo} </p>
                    <Logout />
                </div>
            </nav>

        ) : (
            <ul className='container-fluid list-unstyled nav nav-tabs'>
                <li className=''>
                    <button className="nav-link text-start">{toHome()}</button>
                </li>
            </ul>
        )
    );
}

export default Navbar;