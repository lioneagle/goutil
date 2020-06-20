/* generate switch, left brace at next line */
switch (state)
{
    case ST_INIT:
    {
        /* state1 */
        x += 5;
        break;
    }
    case ST_1:
    {
        /* state2 */
        x += 6;
        break;
    }
    case ST_2:
    {
        /* state3 */
        x += 7;
        break;
    }
}
/* generate switch, left brace at next line */
switch (state)
{
    case ST_INIT:
    {
        /* state1 */
        x += 5;
        break;
    }
    case ST_1:
    {
        /* state2 */
        x += 6;
        break;
    }
    case ST_2:
    {
        /* state3 */
        x += 7;
        break;
    }
    default:
    {
        x -= 3;
        break;
    }
}

/* generate switch, left brace at same line */
switch () {
    case ST_INIT: {
        /* state4 */
        y += 5;
        break;
    }
    case ST_1: {
        /* state5 */
        y += 6;
        break;
    }
    case ST_2: {
        /* state6 */
        y += 7;
        break;
    }
}
/* generate switch, left brace at same line */
switch () {
    case ST_INIT: {
        /* state4 */
        y += 5;
        break;
    }
    case ST_1: {
        /* state5 */
        y += 6;
        break;
    }
    case ST_2: {
        /* state6 */
        y += 7;
        break;
    }
    default: {
        y -= 3;
        break;
    }
}
