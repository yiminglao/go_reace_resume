import React, { Fragment } from 'react';
import { Container, TextField, createStyles, makeStyles, Theme, Button } from '@material-ui/core';
import Title from '../common/Title';
import SendIcon from '@material-ui/icons/Send';

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		inputField: {
			marginTop: 8,
			marginBottom: 8,
			fontSize: '12px'
		},
		button: {
			margin: theme.spacing(1)
		}
	})
);

const ContactFrom = () => {
	const classes = useStyles();
	return (
		<Fragment>
			<Title value='Contact Me' />
			<Container maxWidth='lg'>
				<TextField className={classes.inputField} fullWidth size='small' label='Name' defaultValue='Hello World' variant='outlined' />
				<TextField className={classes.inputField} fullWidth size='small' label='Email' defaultValue='Hello World' variant='outlined' />
				<TextField className={classes.inputField} fullWidth size='small' label='Subject' defaultValue='Hello World' variant='outlined' />
				<TextField className={classes.inputField} fullWidth size='small' label='Message' multiline rows='4' defaultValue='Default Value' variant='outlined' />
				<Button variant='contained' color='default' className={classes.button} endIcon={<SendIcon />}>
					Send
				</Button>
			</Container>
		</Fragment>
	);
};

export default ContactFrom;
