import React from "react"

const Result =
    (props) => {
      return (
          <a className =
               "flex md:text-black text-white bg-gray-200 md:bg-gray-300 w-full z-50 mb-1 md:my-2 py-1 px-2 rounded hover:bg-gray-300 md:hover:bg-gray-400" href =
                   {props.result.url}>
          <span>{props.result.name}</span>
        </a>)
    }

export default Result