import React, {useEffect, useState} from "react"
import {Col, Container, Row} from "react-bootstrap";
import axios from "axios";
import {Link, useLocation} from "react-router-dom";

import {Spinner} from "react-bootstrap";

const Home = () => {
	const [apiData, setApiData] = useState(false);
	const [loading, setLoading] = useState(true);
	
	const location = useLocation();
	
	useEffect(() => {
		const fetchData = async () => {
			
			try {
				const apiUrl = process.env.REACT_APP_API_ROOT;
				const response = await axios.get(apiUrl)
				
				if (response.status === 200) {
					if (response?.data.statusText === "Ok") {
						setApiData(response?.data?.blog_records);
					}
				}
				
				setLoading(false);
			} catch (error) {
				console.log(error.response);
			}
			
		};
		
		fetchData();
		return () => {}
	}, []);
	
	console.log(apiData);
	
	if (loading) {
		return (
			<>
				<Container className="spinner">
					<Spinner animation="grow" />
				</Container>
			</>
		);
	}
	
	return (
		<Container>
			<Row>
				<h3>
					<Link to="add" className="btn btn-primary">
						Add New
					</Link>
				</h3>
				
				<h5>{location.state && location.state}</h5>
				
				{apiData && (
					apiData.map((record, index) => (
							<Col key={index} xs="4" className="py-5 box">
								<div>
									<img src={`${process.env.REACT_APP_API_ROOT}/${record.image}`}/>
								</div>
								<div className="title">
									<Link to={`blog/${record.id}`}> {record.title}</Link>
								</div>
								<div>
									<Link to={`edit/${record.id}`}>
										<i className="fa fa-solid fa-pencil fa-2x"/>
									</Link>&nbsp;
									<Link to={`delete/${record.id}`}>
										<i className="fa fa-solid fa-trash fa-2x"/>
									</Link>
								</div>
								<div>{record.desc}</div>
							</Col>
					))
				)}
			</Row>
		</Container>
	)
}

export default Home;