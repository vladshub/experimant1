import React, {Component} from 'react';
import {Link, Route, Switch} from 'react-router-dom'
import {Avatar, FontIcon, List, ListItem, Snackbar, Subheader} from 'react-md';
import ShowSite from "./ShowSite";
import SitesAPI from "./api";

class ListSites extends Component {
    state = {
        sites: {},
        toasts: [],
        autohide: true
    };

    dismissToast = () => {
        const [, ...toasts] = this.state.toasts;
        this.setState({toasts});
    };

    componentDidMount() {
        this._asyncRequest = SitesAPI.all().then(sites => {
            this.setState({
                sites: sites,
            });
        }).catch((err) => {
            this.setState((state) => {
                const toasts = state.toasts.slice();
                toasts.push({text: "Had an error retriving data", action: false});
                return {toasts};
            });
        });
    }

    componentWillUnmount() {
        if (this._asyncRequest) {
            this._asyncRequest = null;
        }
    }

    render() {
        const {toasts, autohide} = this.state;
        return (
            <List className="md-cell md-cell--12">
                <Subheader primaryText="Sites List"/>
                {
                    Object.keys(this.state.sites).map((i) => {
                        let s = this.state.sites[i];
                        return <ListItem
                            leftAvatar={<Avatar icon={<FontIcon>folder</FontIcon>}/>}
                            rightIcon={<FontIcon>info</FontIcon>}
                            primaryText={s.url}
                            secondaryText={s.id}
                            key={s.id}
                            to={`/sites/${s.id}`}
                            component={Link}
                        />
                    })
                }
                <Switch>
                    <Route path='/sites/:number' component={ShowSite}/>
                </Switch>
                <Snackbar
                    id="interactive-snackbar"
                    toasts={toasts}
                    autohide={autohide}
                    onDismiss={this.dismissToast}
                />
            </List>
        );
    }
}

export default ListSites;