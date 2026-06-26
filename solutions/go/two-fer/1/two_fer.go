/* This package is about a bakery's two-for-one cookie holiday offer. As you take the offer, and decide to offer the extra cookie to someone else in the queue, you can use the ShareWith function to determine what you will say as you do so.
*/
package twofer

// The ShareWith function takes a string ("name") as a parameter and returns another string acording to
// the value of "name"
func ShareWith(name string) string {
	if name == ""{
        return "One for you, one for me."
    } else {
        return "One for " + name + ", one for me." 
    }
}
