import React, {Component} from 'react';
import {Button, Card, CardText, CardTitle, Snackbar, TextField} from 'react-md';
import SitesAPI from "./api";

class AddSite extends Component {
    dismissToast = () => {
        const [, ...toasts] = this.state.toasts;
        this.setState({toasts});
    };

    constructor(props) {
        super(props);
        this.state = {
            toasts: [],
            autohide: true
        };

        // This binding is necessary to make `this` work in the callback
        this.addSite = this.addSite.bind(this);
    };

    addSite(e) {
        e.preventDefault();
        const data = new FormData(e.target);
        var url = data.get('url');
        SitesAPI.add(url).then(() => {
            this.setState((state) => {
                const toasts = state.toasts.slice();
                toasts.push({text: "Successfully added new url", action: false});
                return {toasts};
            });
        }).catch((err) => {
            this.setState((state) => {
                const toasts = state.toasts.slice();
                toasts.push({text: "Got an error from the server please try again later", action: false});
                return {toasts};
            });
        });
    };

    render() {
        const {toasts, autohide} = this.state;
        return (
            <Card className="md-cell md-cell--12 md-text-container">
                <CardTitle title="Inbox"/>
                <CardText>
                    <form
                        id="add_site"
                        ref={this.setForm}
                        onSubmit={this.addSite}
                        name="add_site"
                    >
                        <TextField
                            id="url"
                            name="url"
                            label="Url"
                            placeholder="https://www.google.com"
                            className="md-cell md-cell--top md-cell--12"
                            helpOnFocus
                            helpText="A url that will be crawled"
                        />
                        <Button type="submit" raised primary swapTheming>Add</Button>
                    </form>
                </CardText>
                <Snackbar
                    id="interactive-snackbar"
                    toasts={toasts}
                    autohide={autohide}
                    onDismiss={this.dismissToast}
                />
            </Card>
        )
    };
}

export default AddSite;