import React from 'react';
import style from './action-buttons.module.scss';
import cx from 'classnames';
import buttonsConf from './action-buttons-conf';

interface ButtonConf {
    id: number;
    icon: string;
    label: string;
    style: string
}

function ActionButtons() {
    
    const displayButtons = buttonsConf.map((button: ButtonConf) => {
        return (
            <button key={button.id} className={cx(style.actionButton, style[button.style])}>
                <i className={cx(button.icon, style.buttonIcon)}></i>
                <span className={style.buttonLabel}>{button.label}</span>
            </button>
        )
    });

    return (
        <div>
            <div className={style.buttonsContainer}>
                {displayButtons}
            </div>
        </div>
    )
}

export default ActionButtons;
