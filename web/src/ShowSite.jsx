import React, {Component} from 'react';
import {Avatar, FontIcon, List, ListItem, Snackbar, Subheader} from 'react-md';
import PropTypes from 'prop-types';
import SitesAPI from "./api";

class ShowSite extends Component {
    static propTypes = {
        location: PropTypes.object.isRequired
    };
    state = {
        siteId: "",
        site: {},
        fbReachEstimate: {},
        toasts: [],
        autohide: true
    };

    //
    //     const site = SitesAPI.get(parseInt(props.match.params.number, 10))
    // if (!site) {
    //     return <div>Site Not Found</div>
    dismissToast = () => {
        const [, ...toasts] = this.state.toasts;
        this.setState({toasts});
    };

    // }
    constructor(props) {
        super(props);
        this.state.siteId = props.match.params.number;

    };

    componentDidUpdate(prevProps) {
        if (this.props.location !== prevProps.location) {
            this.updateData();
        }
    }

    updateData() {
        this._asyncRequest = SitesAPI.get(this.state.siteId).then(site => {
            debugger;
            this.setState({
                fbReachEstimate: { audianceSize: site.audianceSize, estimateReady: site.estimateReady },
                site: site,
                siteId: this.props.match.params.number
            });
        }).catch((err) => {
            this.setState((state) => {
                const toasts = state.toasts.slice();
                toasts.push({text: "Had an error retriving data", action: false});
                return {toasts};
            });
        });

    }

    componentDidMount() {
        this.updateData()
    }

    componentWillUnmount() {
        if (this._asyncRequest) {
            this._asyncRequest = null;
        }
    }

    render() {
        const {toasts, autohide} = this.state;
        debugger;
        return (
            <List className="md-cell md-cell--12">
                <Subheader primaryText={this.state.siteId}/>
                    <ListItem
                        key="1"
                        leftAvatar={<Avatar icon={<FontIcon>insert_chart</FontIcon>}/>}
                        primaryText={"audianceSize:" + this.state.fbReachEstimate.audianceSize}
                        secondaryText={"estimateReady:" + this.state.fbReachEstimate.estimateReady}
                    />


                <Snackbar
                    id="interactive-snackbar"
                    toasts={toasts}
                    autohide={autohide}
                    onDismiss={this.dismissToast}
                />
            </List>
        )

    };
};
export default ShowSite;