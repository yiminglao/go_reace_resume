import React, { Fragment } from 'react';
import { Typography, Container } from '@material-ui/core';
import Title from '../common/Title';

const AboutMe = () => {
	return (
		<Fragment>
			<Title value='About Me' subTitle='Software developer' />
			<Container maxWidth='lg' style={{ height: 300 }}>
				<Typography variant='body1' component='p' align='left'>
					At Helio Training, we offer bootcamps which prepare you to work in the tech industry. <br /> You’ll learn in an intense, tech-focused environment from instructors who are passionate educators with years of experience in the tech
					industry. You’ll learn leading-edge development skills while you work on real-world projects with industry partners such as the Utah Jazz and Larry H. Miller Sports and Entertainment. Best of all, you’ll graduate with the skills
					you need to launch your new career. At Helio Training, we offer bootcamps
				</Typography>
			</Container>
		</Fragment>
	);
};

export default AboutMe;
