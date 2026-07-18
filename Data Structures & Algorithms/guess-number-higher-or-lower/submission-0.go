/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l,r := 0, n
	for {
		m := l + (r-l)/2
		target := guess(m)
		if target == 0{
			return m
		} else if target>0{
			l=m+1
		} else {
			r = m-1
		}
	}
}
