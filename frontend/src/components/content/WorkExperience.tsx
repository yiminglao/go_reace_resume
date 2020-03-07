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

const WorkExperience = () => {
	const classes = useStyles();
	return (
		<Fragment>
			<Title value='Work Experience' />
			<Container maxWidth='lg' style={{ height: 300 }}>
				<List className={classes.root}>
					<ListItem alignItems='flex-start'>
						<ListItemAvatar>
							<Avatar alt='bc logo' src='./images/bc_logo.png' />
						</ListItemAvatar>
						<ListItemText
							primary='Basecamp Franchising '
							secondary={
								<React.Fragment>
									<Typography component='span' variant='body2' className={classes.inline} color='textPrimary'>
										Software developer — Salt Lake City ​
									</Typography>
								</React.Fragment>
							}
						/>
					</ListItem>
				</List>
			</Container>
		</Fragment>
	);
};

export default WorkExperience;
