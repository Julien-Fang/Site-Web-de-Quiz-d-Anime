// Menu.jsx
import { useContext, useState, useEffect } from 'react';
import { QuizContext, UserContext } from '../AppContext';
import Answer from './Answer';
import axios from 'axios';


const Quiz = () => {
	const { quizData, num } = useContext(QuizContext);
	const { handleLogout } = useContext(UserContext);

	const [pinned,setPinned] = useState(false);//Si le post est épinglé
	const [refreshPinned,setRefreshPinned] = useState(false);//Actualiser l'état du pin

	const questionFormat = (question) => {
		var parts = question.split(': ');
		const firstPart = parts.shift(); // Récupère la première partie
		const remainingString = parts.join(': '); // Joindre les parties restantes avec ':'
		parts = [firstPart, remainingString]
		
		const formattedText = []
		parts.map((part, index) => {
		  formattedText.push(<span key={index}>{part}</span>);
		});
	  
	  	return formattedText;
	};

	useEffect(() => {
		if (num === quizData.quizcollection.length) {
			window.location.href = "http://localhost:3000";
		}
		console.log(quizData)

	}, [num, quizData])



	useEffect(() => {
		axios({
			method : 'GET',
			url : `http://localhost:8000/pin/find`,
			withCredentials : true
		})
		.then (res => {
			if (res.data.includes(quizData.quizcollection[num].anime)) {
				setPinned(pinned);
			}
			else {
				setPinned(false);
			}
		})
		.catch(err => {
			handleLogout();
		})
	}, [refreshPinned, num, pinned, quizData.quizcollection])

		

	const handlePin = () => {
		if(pinned){
			axios({
				method : 'DELETE',
				url : `http://localhost:8000/pin/delete`,
				data : {
					anime : quizData.quizcollection[num].anime
				},
				withCredentials : true
			})
			.then (res => {
				setPinned(false);
				setRefreshPinned(!refreshPinned);
			}
			)
			.catch(err => {
				console.log(err);
			}
			)
		}else{
			axios({
				method : 'POST',
				url : `http://localhost:8000/pin/create`,
				data : {
					anime : quizData.quizcollection[num].anime
				},
				withCredentials : true
			})
			.then (res => {
				setPinned(true);
				setRefreshPinned(!refreshPinned);
			}
			)
			.catch(err => {
				console.log(err);
			}
			)
		}
	}



	return (
        <div className='quiz col-md-8'>
            {quizData && num < 10 && (
                <div className="container">

                    <div className="row justify-content-center">
						<h1 className='text-light col-4 my-5'>Question {num + 1}</h1>
                        <div className="col-md-10 mt-1">
                            <div className="progress">
                                <div className="progress-bar progress-bar-striped bg-warning progress-bar-animated" role="progressbar" style={{ width: `${num * 10}%` }} aria-valuenow={num * 10} aria-valuemin="0" aria-valuemax="100">{num * 10}%</div>
                            </div>
                        </div>

						<div className="col-10 mt-4" style={{backgroundColor: '#383838'}}>
							<button id="pin" type="button" style={ pinned ? { backgroundColor:"lightgreen"} : {backgroundColor :"lightgray"}} onClick={handlePin} >&#x1F4CC;</button>
							<p style={{fontWeight: "bold", color: "white"}}>{questionFormat(quizData.quizcollection[num].question)}</p>

							{quizData.quizcollection[num].image && (
								<div className="d-flex justify-content-center align-items-center my-4" style={{ height: '300px', overflow: 'hidden' }}>
									<img
										src={quizData.quizcollection[num].image}
										alt="Anime"
										className="img-fluid"
										style={{ width: 'auto', height: '100%', objectFit: 'contain' }}
									/>
								</div>
							)}

							<Answer num={num}/>
						</div>
                    </div>
                </div>
            )}
        </div>
    );
}

export default Quiz;
