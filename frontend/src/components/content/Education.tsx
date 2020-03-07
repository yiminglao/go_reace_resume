import React, { Fragment } from 'react';
import { Container, createStyles, makeStyles, Theme, List, ListItem, ListItemAvatar, Avatar, ListItemText, Typography, Divider } from '@material-ui/core';
import Title from '../common/Title';

const useStyles = makeStyles(() =>
	createStyles({
		root: {
			width: '100%'
		},
		inline: {
			display: 'inline'
		}
	})
);

const Education = () => {
	const classes = useStyles();
	return (
		<Fragment>
			<Title value='EDUCATION' />
			<Container maxWidth='lg' style={{ height: 300 }}>
				<List className={classes.root}>
					<ListItem alignItems='flex-start'>
						<ListItemAvatar>
							<Avatar alt='Remy Sharp' src='./images/weber_logo.png' />
						</ListItemAvatar>
						<ListItemText
							primary='Weber State University'
							secondary={
								<React.Fragment>
									<Typography component='span' variant='body2' className={classes.inline} color='textPrimary'>
										Bachelor in Computer Science
									</Typography>
									{' — 08/2015 - 12/2018'}
								</React.Fragment>
							}
						/>
					</ListItem>
					<Divider variant='inset' component='li' />
					<ListItem alignItems='flex-start'>
						<ListItemAvatar>
							<Avatar alt='udemy' src='./images/udemy_logo.png' />
						</ListItemAvatar>
						<ListItemText
							primary='udemy'
							secondary={
								<React.Fragment>
									<Typography component='span' variant='body2' className={classes.inline} color='textPrimary'>
										Self Study
									</Typography>
								</React.Fragment>
							}
						/>
					</ListItem>
					<Divider variant='inset' component='li' />
				</List>
			</Container>
		</Fragment>
	);
};

export default Education;
