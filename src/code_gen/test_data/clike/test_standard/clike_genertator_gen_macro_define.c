/* NAME: parse
 * PARAMS:
 * context -- context for parsing
 * src     -- source to parse
 * len     -- length of source
 * RETURN:
 * length parsed
 * NOTE: create for test
 */
#define PARSE(context, src, len) do{return 0;}while(0)
/*********************************
 * NAME: parse
 * PARAMS:
 * context -- context for parsing
 * src     -- source to parse
 * len     -- length of source
 * RETURN:
 * length parsed
 * NOTE: create for test
 **********************************/
#define PARSE(context, src, len) \
do{\
    x = 1;\
    y = 2;\
    return 1;\
}while(0);