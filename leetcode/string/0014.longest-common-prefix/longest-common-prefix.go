

// 1.描述
/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

链接：https://leetcode-cn.com/problems/longest-common-prefix

*/

/*
最长公共前缀

方法一：水平扫描法
思路

首先，我们将描述一种查找一组字符串的最长公共前缀 LCP(S1...Sn)的简单方法。
我们将会用到这样的结论：

LCP(S1...Sn) = LCP(LCP(LCP(S1, S2),S3),...Sn)

算法

为了运用这种思想，算法要依次遍历字符串 [S1...Sn]，当遍历到第 i个字符串的时候，找到最长公共前缀 LCP(S1...Si)。
当 LCP(S1...Si)是一个空串的时候，算法就结束了。 否则，在执行了 n次遍历之后，算法就会返回最终答案 LCP(S1...Sn)。

找到最长公共前缀

查找最长公共前缀 （水平扫描法）
golang版代码：
func longestCommonPrefix(strs []string) string {
    if(len(strs)) == 0{
        return ""
    }
    prefix := strs[0]
    for i := 1; i < len(strs); i++{                // 追个字符串比较，看看是否有共同的前缀prefix
        for strings.Index(strs[i], prefix) != 0 {  // 如果当前子串prefix不是字符串strs[i]的前缀（前几个字符相同），则对前缀进行更新（去掉末尾字符），直至前缀为空返回空字符串
            prefix = prefix[:len(prefix) - 1]      // 更新前缀（去掉最后一个字符）
            if(len(prefix) == 0 ){
                return ""
            }
        }
    }
    return prefix
}




java版代码：
public String longestCommonPrefix(String[] strs) {
   if (strs.length == 0) return "";
   String prefix = strs[0];
   for (int i = 1; i < strs.length; i++)
       while (strs[i].indexOf(prefix) != 0) {
           prefix = prefix.substring(0, prefix.length() - 1);
           if (prefix.isEmpty()) return "";
       }        
   return prefix;
}

复杂度分析

时间复杂度：O(S)，S 是所有字符串中字符数量的总和。

最坏的情况下，n个字符串都是相同的。算法会将 S1 与其他字符串 [S2... Sn]都做一次比较。
这样就会进行 S次字符比较，其中 S 是输入数据中所有字符的数量。

空间复杂度：O(1)，我们只需要使用常数级别的额外空间。

算法二：水平扫描
算法

想象数组的末尾有一个非常短的字符串，使用上述方法依旧会进行 S​ 次比较。优化这类情况的一种方法
就是水平扫描。从前往后枚举字符串的每一列，先比较每个字符串相同列上的字符（即不同字符串
相同下标的字符）然后再进行下一列的比较。
golang版代码：
func longestCommonPrefix(strs []string)string{
    if(len(strs)) == 0 {
        return ""
    }

    prefix := strs[0]
    for i := 0; i < len(prefix); i++{   // 在strs[0]长度范围内进行比较
        currentChar := prefix[i]       // 将strs[0]字符串的逐个元素与其他串比较
        for j := 1; j < len(strs); j++{
            if(len(strs[j]) == i || strs[j][i] != currentChar){  // 如果当前串的长度在strs[0]范围内，或者其第i个字符与strs[0]中对应位置字符不一致，其公共前缀prefix就更新（截取0至i-1索引处的字符返回）
                return prefix[:i]   
            }
        }
    }
    return prefix  // 返回strs[0]
}

java版代码：
public String longestCommonPrefix(String[] strs) {
    if (strs == null || strs.length == 0) return "";
    for (int i = 0; i < strs[0].length() ; i++){
        char c = strs[0].charAt(i);
        for (int j = 1; j < strs.length; j ++) {
            if (i == strs[j].length() || strs[j].charAt(i) != c)
                return strs[0].substring(0, i);             
        }
    }
    return strs[0];
}
复杂度分析

时间复杂度：O(S)，S 是所有字符串中字符数量的总和。

最坏情况下，输入数据为 n 个长度为 m 的相同字符串，算法会进行 S = m*n 次比较。
可以看到最坏情况下，本算法的效率与算法一相同，但是最好的情况下，算法只需要进行 n*minLen 次比较，其中 minLen 是数组中最短字符串的长度。

空间复杂度：O(1)，我们只需要使用常数级别的额外空间。

算法三：分治
思路

这个算法的思路来自于LCP操作的结合律。 我们可以发现：
LCP(S1...Sn) = LCP(LCP(S1...Sk), LCP (S{k+1}...Sn))，其中 LCP(S1...Sn)
是字符串 [S1...Sn]的最长公共前缀，1 < k < n。

算法
为了应用上述的结论，我们使用分治的技巧，将原问题 LCP(Si...Sj)
分成两个子问题 LCP(Si...S{mid})与 LCP(S{mid+1}...Sj)，其中 
mid =(i+j)/2。 我们用子问题的解 lcpLeft 与 lcpRight 构造原问题的解 LCP(Si...Sj)。 
从头到尾挨个比较 lcpLeft 与 lcpRight 中的字符，直到不能再匹配为止。 
计算所得的 lcpLeft 与 lcpRight 最长公共前缀就是原问题的解 LCP(Si...Sj)。


寻找最长公共前缀的分治方法
golang代码：
// 分治法
func longestCommonPrefix(strs []string)string{
    if(len(strs)) == 0 {
        return ""
    }
    return longestDividedCommonPrefix(strs, 0, len(strs) -1)
}

func longestDividedCommonPrefix(strs []string , left, right int)string {
    if(left == right){
        return strs[left]
    }else{
        mid := (left + right) / 2
        lcpLeft := longestDividedCommonPrefix(strs, left, mid)
        lcpRight := longestDividedCommonPrefix(strs, mid+1, right)
        return commonPrefix(lcpLeft, lcpRight)
    }
}

func commonPrefix(left , right string)string{
    min := minInt(len(left), len(right))
    for i := 0; i < min; i++ {
        if(left[i] != right[i]){
            return left[:i]
        }
    }
    return left[:min]
}

func minInt(l, r int)int{
    if(l < r){
        return l
    }else{
        return r
    }
}



java版代码：
public String longestCommonPrefix(String[] strs) {
    if (strs == null || strs.length == 0) return "";    
        return longestCommonPrefix(strs, 0 , strs.length - 1);
}

private String longestCommonPrefix(String[] strs, int l, int r) {
    if (l == r) {
        return strs[l];
    }
    else {
        int mid = (l + r)/2;
        String lcpLeft =   longestCommonPrefix(strs, l , mid);
        String lcpRight =  longestCommonPrefix(strs, mid + 1,r);
        return commonPrefix(lcpLeft, lcpRight);
   }
}

String commonPrefix(String left,String right) {
    int min = Math.min(left.length(), right.length());       
    for (int i = 0; i < min; i++) {
        if ( left.charAt(i) != right.charAt(i) )
            return left.substring(0, i);
    }
    return left.substring(0, min);
}
复杂度分析

最坏情况下，我们有 nn 个长度为 mm 的相同字符串。

时间复杂度：O(S)O(S)，SS 是所有字符串中字符数量的总和，S=m*nS=m∗n。

时间复杂度的递推式为 T(n)=2\cdot T(\frac{n}{2})+O(m)T(n)=2⋅T( 
2
n
​	
 )+O(m)， 化简后可知其就是 O(S)O(S)。最好情况下，算法会进行 minLen\cdot nminLen⋅n 次比较，其中 minLenminLen 是数组中最短字符串的长度。

空间复杂度：O(m \cdot log(n))O(m⋅log(n))

内存开支主要是递归过程中使用的栈空间所消耗的。 一共会进行 log(n)log(n) 次递归，每次需要 mm 的空间存储返回结果，所以空间复杂度为 O(m\cdot log(n))O(m⋅log(n))。

方法四：二分查找法
这个想法是应用二分查找法找到所有字符串的公共前缀的最大长度 L。 算法的查找区间是 (0 \ldots minLen)(0…minLen)，其中 minLen 是输入数据中最短的字符串的长度，同时也是答案的最长可能长度。 每一次将查找区间一分为二，然后丢弃一定不包含最终答案的那一个。算法进行的过程中一共会出现两种可能情况：

S[1...mid] 不是所有串的公共前缀。 这表明对于所有的 j > i S[1..j] 也不是公共前缀，于是我们就可以丢弃后半个查找区间。

S[1...mid] 是所有串的公共前缀。 这表示对于所有的 i < j S[1..i] 都是可行的公共前缀，因为我们要找最长的公共前缀，所以我们可以把前半个查找区间丢弃。

使用二分法寻找最长公共前缀

图 3. 使用二分查找法寻找最长公共前缀

public String longestCommonPrefix(String[] strs) {
    if (strs == null || strs.length == 0)
        return "";
    int minLen = Integer.MAX_VALUE;
    for (String str : strs)
        minLen = Math.min(minLen, str.length());
    int low = 1;
    int high = minLen;
    while (low <= high) {
        int middle = (low + high) / 2;
        if (isCommonPrefix(strs, middle))
            low = middle + 1;
        else
            high = middle - 1;
    }
    return strs[0].substring(0, (low + high) / 2);
}

private boolean isCommonPrefix(String[] strs, int len){
    String str1 = strs[0].substring(0,len);
    for (int i = 1; i < strs.length; i++)
        if (!strs[i].startsWith(str1))
            return false;
    return true;
}
复杂度分析

最坏情况下，我们有 nn 个长度为 mm 的相同字符串。

时间复杂度：O(S \cdot log(n))O(S⋅log(n))，其中 SS 所有字符串中字符数量的总和。

算法一共会进行 log(n)log(n) 次迭代，每次一都会进行 S = m*nS=m∗n 次比较，所以总时间复杂度为 O(S \cdot log(n))O(S⋅log(n))。

空间复杂度：O(1)O(1)，我们只需要使用常数级别的额外空间。

更进一步
让我们看一个有些不同的问题：

给定一些键值字符串 S = [S_1,S_2 \ldots S_n][S 
1
​	
 ,S 
2
​	
 …S 
n
​	
 ]，我们要找到字符串 q 与 S 的最长公共前缀。 这样的查询操作可能会非常频繁。

我们可以通过将所有的键值 S 存储到一颗字典树中来优化最长公共前缀查询操作。 如果你想学习更多关于字典树的内容，可以从 208. 实现 Trie (前缀树) 开始。在字典树中，从根向下的每一个节点都代表一些键值的公共前缀。 但是我们需要找到字符串q 和所有键值字符串的最长公共前缀。 这意味着我们需要从根找到一条最深的路径，满足以下条件：

这是所查询的字符串 q 的一个前缀

路径上的每一个节点都有且仅有一个孩子。 否则，找到的路径就不是所有字符串的公共前缀

路径不包含被标记成某一个键值字符串结尾的节点。 因为最长公共前缀不可能比某个字符串本身长

算法

最后的问题就是如何找到字典树中满足上述所有要求的最深节点。 最有效的方法就是建立一颗包含字符串 [S_1 \ldots S_n][S 
1
​	
 …S 
n
​	
 ] 的字典树。 然后在这颗树中匹配 q 的前缀。 我们从根节点遍历这颗字典树，直到因为不能满足某个条件而不能再遍历为止。

使用字典树寻找最长公共前缀

图 4. 使用字典树查找最长公共前缀

public String longestCommonPrefix(String q, String[] strs) {
    if (strs == null || strs.length == 0)
         return "";  
    if (strs.length == 1)
         return strs[0];
    Trie trie = new Trie();      
    for (int i = 1; i < strs.length ; i++) {
        trie.insert(strs[i]);
    }
    return trie.searchLongestPrefix(q);
}

class TrieNode {

    // 子节点的链接数组
    private TrieNode[] links;

    private final int R = 26;

    private boolean isEnd;

    // 非空子节点的数量
    private int size;    
    public void put(char ch, TrieNode node) {
        links[ch -'a'] = node;
        size++;
    }

    public int getLinks() {
        return size;
    }
    // 假设方法 containsKey、isEnd、get、put 都已经实现了
    // 可以参考文章：https://leetcode.com/articles/implement-trie-prefix-tree/
}

public class Trie {

    private TrieNode root;

    public Trie() {
        root = new TrieNode();
    }

// 假设方法 insert、search、searchPrefix 都已经实现了
// 可以参考文章：https://leetcode.com/articles/implement-trie-prefix-tree/
    private String searchLongestPrefix(String word) {
        TrieNode node = root;
        StringBuilder prefix = new StringBuilder();
        for (int i = 0; i < word.length(); i++) {
            char curLetter = word.charAt(i);
            if (node.containsKey(curLetter) && (node.getLinks() == 1) && (!node.isEnd())) {
                prefix.append(curLetter);
                node = node.get(curLetter);
            }
            else
                return prefix.toString();

         }
         return prefix.toString();
    }
}
复杂度分析

最坏情况下查询字符串 qq 的长度为 mm 并且它与数组中 nn 个字符串均相同。

时间复杂度：预处理过程 O(S)O(S)，其中 SS 数组里所有字符串中字符数量的总和，最长公共前缀查询操作的复杂度为 O(m)O(m)。

建立字典树的时间复杂度为 O(S)O(S)。在字典树中查找字符串 qq 的最长公共前缀在最坏情况下需要 O(m)O(m) 的时间。

空间复杂度：O(S)O(S)，我们只需要使用额外的 SS 空间建立字典树。
*/