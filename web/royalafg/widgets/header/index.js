import {faBars, faTimes} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {useEffect, useState} from "react";
import React from "react"

import IdNav from "./idnav";
import NavItems from "./nav";

class Header extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      isOpen : false,
      width : 0
    }

                 this.updateDimensions = this.updateDimensions.bind(this);
    this.toggleMenu = this.toggleMenu.bind(this);
  }

  toggleMenu() { this.setState(state => ({isOpen : !state.isOpen})); }

  updateDimensions = () => { this.setState({width : window.innerWidth}); };
  componentDidMount() {
    this.updateDimensions();
    window.addEventListener('resize', this.updateDimensions);
  }
  componentWillUnmount() {
    window.removeEventListener('resize', this.updateDimensions);
  }

  render() {
    return (
        <div><header className = "md:h-10 h-16 bg-blue-600 text-white">
        <div className = "flex h-full">
        <button className =
             "hamburger md:hidden flex h-full items-center ml-6 z-10 cursor-pointer focus:outline-none w-6" onClick =
                 {this.toggleMenu}>
        <FontAwesomeIcon icon = {this.state.isOpen ? faTimes : faBars} size =
             "lg" />
        </button>
                        <div className="logo md:ml-16 flex items-center md:h-full h-16 w-full md:w-auto justify-center md:relative absolute mt-0 mb-auto md:z-10 z-0">
                            <a className="font-medium font-sans text-xl text-center cursor-pointer" href="/">Royalafg</a>
         <
         /div>
                        {
                            this.state.width <= 768 && !this.state.isOpen ? <></>:
            <div className =
                 "nav md:ml-4 md:flex absolute md:relative w-full md:h-full z-50 md:z-10 bg-gray-200 md:bg-transparent text-black md:text-white mt-16 md:mt-0">
        <NavItems /></div>
                        }

                    </div>
        </header>
            </div>)
  }
}

export default Header