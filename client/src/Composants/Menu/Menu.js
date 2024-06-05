// Menu.jsx
import { useEffect, useContext, useState } from 'react';
import axios from 'axios';

import { QuizContext, UserContext } from '../AppContext';
import MenuItem from './MenuItem';
import Trend from './Trend';
import Recent from './Recent';

const Menu = () => {
	const { quizType, setQuizData } = useContext(QuizContext);
	const { handleLogout } = useContext(UserContext);
	const [unfinishedQuiz, setUnfinishedQuiz] = useState(false);

	useEffect(() => {
        if (quizType !== "") {
			axios({
				method: 'POST',
				url: `http://localhost:8000/anime/${quizType}`,
				withCredentials: true,
			})
			.then((res) => {
				setQuizData(res.data);
			})
			.catch(err => {
				handleLogout();
			});
        }

		axios({
            method : 'GET',
            url : `http://localhost:8000/quiz/unfinishedQuiz`,
            withCredentials : true
        })
        .then (res => {
            setUnfinishedQuiz(res.data);
        })
		.catch(err => {
			handleLogout();
		})
    }, [quizType, setQuizData, handleLogout])

	return (
		<div className="menu row" style={{ height: '100vh', overflow: 'auto' }}>

			<div className="menu-container col-md-8 d-flex flex-column align-items-center justify-content-center">
		
				<div className="col-md-6">
						<MenuItem label="General" />
						<MenuItem label="Genre" />
						<MenuItem label="Picture" />
						<MenuItem label="Synopsis" />
				</div>

				{unfinishedQuiz && unfinishedQuiz.length > 0 ? 
					<div className="m-5" style={{color: 'red', fontSize:'25px'}}>Quiz à terminer avant de consulter les pins:<br/>
						{unfinishedQuiz.map((quiz, numero) => (
							<span key={"unfinishedQuiz"+numero}> {quiz} </span>
						))}
					</div> : null}
			</div>

			<div className="table-container col-md-4">
				<Trend />
				<Recent />
			</div>
		</div>
		
	);
}

export default Menu;
