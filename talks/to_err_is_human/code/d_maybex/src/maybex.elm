module Maybex exposing (main)

import Html exposing (text)


content num =
    -- BGN OMIT
    case String.toInt num of
        Just n ->
            "The integer is " ++ String.fromInt n ++ "."

        Nothing ->
            "The value is not a valid integer."



-- END OMIT


main =
    text <| content ".333"
