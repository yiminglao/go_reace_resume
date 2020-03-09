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

const GitInfo = () => {
	const classes = useStyles();
	return (
		<Fragment>
			<Title value='Projects' />
			<Container maxWidth='lg' style={{ height: 300 }}>
				<List className={classes.root}>
					<ListItem alignItems='flex-start'>
						<ListItemAvatar>
							<Avatar alt='Remy Sharp' src='/static/images/avatar/1.jpg' />
						</ListItemAvatar>
						<ListItemText
							primary='Brunch this weekend?'
							secondary={
								<React.Fragment>
									<Typography component='span' variant='body2' className={classes.inline} color='textPrimary'>
										Ali Connors
									</Typography>
									{" — I'll be in your neighborhood doing errands this…"}
								</React.Fragment>
							}
						/>
					</ListItem>
					<Divider variant='inset' component='li' />
					<ListItem alignItems='flex-start'>
						<ListItemAvatar>
							<Avatar alt='Remy Sharp' src='/static/images/avatar/1.jpg' />
						</ListItemAvatar>
						<ListItemText
							primary='Brunch this weekend?'
							secondary={
								<React.Fragment>
									<Typography component='span' variant='body2' className={classes.inline} color='textPrimary'>
										Ali Connors
									</Typography>
									{" — I'll be in your neighborhood doing errands this…"}
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

export default GitInfo;
