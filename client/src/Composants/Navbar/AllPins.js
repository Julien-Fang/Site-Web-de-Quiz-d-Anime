import { useState, useEffect } from 'react';
import axios from 'axios';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';


const AllPins = () => {
    const [show, setShow] = useState(false);
    const [pins, setPins] = useState([]);
    const [refresh, setRefresh] = useState(false);

    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);

    useEffect(() => {
        axios({
            method : 'GET',
            url : `http://localhost:8000/pin/getAllPin`,
            withCredentials : true
        })
        .then (res => {
            setPins(res.data);
        })
    }
    , [refresh])

    const handleDeletePin = (anime) => {
        axios({
            method : 'DELETE',
            url : `http://localhost:8000/pin/delete`,
            data : {
                anime : anime
            },
            withCredentials : true
        })
        .then (res => {
            setPins(pins.filter(pin => pin !== anime));
            setRefresh(!refresh);
        })
    }


    return (
        <div className='allPins'>
            <Button variant="primary" 
            className="btn btn-primary d-flex align-items-center justify-content-center"
            style={{ borderRadius: "10px", height: "40px", width: "120px", fontSize: "1.1rem" }} onClick={handleShow}
            >
                <span>Pins</span>
            </Button>

            <Modal show={show} onHide={handleClose}>
                <Modal.Header closeButton>
                    <Modal.Title style={{margin: '0 auto'}}>Pins enregistrés</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    {pins.length === 0 
                        ? (
                            <p style={{color: 'red', fontSize:'25px'}}>Chargement...</p> 
                        ) : (
                            <div className="card-container">
                                <ul className="list-group">
                                    {pins.map((anime, index) => (
                                        <li key={index} className="list-item d-flex justify-content-between align-items-center">
                                            <span className="index">{index + 1}.</span>
                                            <span className="name" style={{marginRight: "1rem"}}>{anime}</span>
                                            <button className="btn btn-danger" onClick={() => handleDeletePin(anime)}>Supprimer</button>
                                        </li>
                                    
                                    ))}
                                </ul>
                            </div>
                        )}
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="primary" onClick={handleClose}>
                        Close
                    </Button>
                </Modal.Footer>
            </Modal>
        </div>
    );

    }

export default AllPins;