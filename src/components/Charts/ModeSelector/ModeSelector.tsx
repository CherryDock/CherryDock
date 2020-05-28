import React, { useState } from 'react'
import style from './mode-selector.module.scss';
import Switch from "react-switch";
import cx from 'classnames';

/*
$button-red: #f2877c;
$button-blue: #6fb2ed;
*/

function ModeSelector() {

    const [historyMode, sethistoryMode] = useState<boolean>(false);

    const historyColor = "#838485"
    const instantColor = "#838485"

    function handleChange(checked: boolean) {
        sethistoryMode(checked);
    }

    const instantLabelStyle = historyMode ? style.labelUnselected : style.labelSelected;
    const historyLabelStyle = historyMode ? style.labelSelected : style.labelUnselected;

    return (
        <div className={style.switchContainer}>
            <div className={style.realTimeContainer}>
                <span className={cx(style.label, instantLabelStyle)}>Real Time</span>
                <i className={cx("fas fa-sync", style.icon)}></i>
            </div>
            <Switch
                className={style.switch}
                onChange={handleChange}
                checked={historyMode}
                height={22}
                width={50}
                checkedIcon={false}
                uncheckedIcon={false}
                onColor={historyColor}
                offColor={instantColor}
            />
            <div className={style.historyContainer}>
                <i className={cx("fas fa-history", style.icon)}></i>
                <span className={cx(style.label, historyLabelStyle)}>History</span>
            </div>
            <button className={style.expandButton}>
                <div className={style.toolTip}>
                    <i className={cx("far fa-question-circle", style.icon)}></i>
                    <span className={style.toolTipText}>Full screen mode</span>
                </div>
                <i className={cx("fas fa-expand", style.icon)}></i>
            </button>
        </div>
    )
}

export default ModeSelector
