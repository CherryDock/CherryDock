import React, { useState, useEffect } from 'react';
import style from './container-details.module.scss';
import { useParams } from "react-router-dom";
import ActionButtons from '../../components/ContainerDetails/ActionButtons/ActionButtons';

function ContainerDetails() {
    const { id } = useParams();

    return (
        <div>
            <div className={style.actionButons}>
            <ActionButtons />
            </div>
        </div>
    )
}

export default ContainerDetails
