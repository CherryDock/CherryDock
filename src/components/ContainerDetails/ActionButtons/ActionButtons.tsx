import React from 'react';
import cx from 'classnames';
import style from './action-buttons.module.scss';

function ActionButtons() {
    return (
        <div className={style.actionButtonsContainer}>
            <button className={style.actionButton}>
                <i className={cx("fas fa-clipboard-list", style.actionIcon)}></i>
                <span className={style.actionLabel}>Logs</span>
            </button>

            <button className={style.actionButton}>
                <i className={cx("fas fa-terminal", style.actionIcon)}></i>
                <span className={style.actionLabel}>Run Bash</span>
            </button>
        </div>
    )
}

export default ActionButtons
