package day06

import (
	"fmt"
)

const input = "czfztznzpnpqpfftcftcchfhjjsqsvvffhwfhhqnhqqmlldpdcdnnfvvmpmzpmmrsrwssqbsbsmmdjdssfqfllclcdczdccppvzvtzznpnfpnnljnnpcphpjprrdhddsgdssbqqghqqmgqgccjffdtdcdmdjdcctltdddvhdhfhnnrffwjfjpjnpnpddhbbrrflrldlqlrlqqwssqsrswrrrmcmrcrfrwfwpfpdfppslsjsjhhjjqhqwwzzhttcvvzrrrnlrlmljlrjrsshwwndnnzwzpzcpcllfrrsjsnjncjcvcllrhrbhhdsdqdpqpffmppdzztggvpvnvwvpvdvqddznzrzppznnpnhhwjjrpjpvvqvpvrvfvjvzjvzjzqzlqqwsqqrsrprqpqmmbhbffbnnlzzhszhznnhzhmzzfhflfccvvbnbtnnlcnndjjcvjjshhssfjfcjjpzpqpmmdqmqmmbbldlwwvdwvddznncwcqwcwdwbdwwgzzzqvqqlqlttbppgtptlldldblbbjcbcttbpbvvhlhggghvvqsshjhjjzqzwwcggsfgsgjjgbjgbbgllhhmqmggmsggzszvvzggdngnznlzlrlprpjpgjgvgppfhhbqqcddvvbhbqqmlmsssbbszzfdzzznssjllwlplttdgtdttsmspmpvpgpspqpdqdppfvpvvmggrwwgswswmmmnppmttqppvttrvrllcmctcqqlvvpwwghhlmhmwhmmpfmfvfwwbgbqqdbqblqqclljcljccnhhghbgbrrnrppchhmpmdmfmvffvqqtnqqfccgcfgcgwgcwggvgzvvbwbgbjbwbcbgbzztbbsbdbvddzrdrjjlqqlslrssjddhbddzbzvbvhhrjhhrtrltljljwljjrnjjjlpjphjphjhjvjgjbbwfwcctgtmtmdttjhthphhchfcfnnqlqhqvvjcclczlzwlzlhzlzdlzlbzbvbmbcmcjcncwncntndtntbbgqqbrbsrssrnnrcrmrlmmrqqddhzdhhnqqvccbmbpbwpbwbrwwwgdgnnhzzhphbhwhwzwhhvhqqnmqnmnjmnjjjjztzvzvbzbbvcvggcdchhhhbfbttmmtsmsgmmqvqdvdrvrqrdqqrhqqpsqqfnntsnnslsggcqqrnrmrfffmgmsmgggrdgdrgddzfzqfqfcqfcfrfwrrrgglwlmlmlgmlmrlljzlzglzztllgnndhdthtwwdrrjnncssdsnszsggdhdpprnnqzzltzlzznlzlflgfgbfggqtggfbbccchdcdrrnwnbnhbnbjjpzjpzzwjwcwggfqfccdzzfddmbddgbgrgllnttrwrjwrrbsrsbbwqwmmlggzqgzqzrqzzffmcmssgjjjcdcssndsdfsfddcffbrrcvccwhhzzlglhhsccvdvjjdldndrrvlrrbwrrbqqcbcjcjpjggngvnndtdvttlrlmmcbcggbtbhthjttdggtvtmmlgmlgmmdwdndmmldlttjlmbtrwfnvmnzggqfqwnppbrnjmpjfppwltpqfstgnnjwmbffgfbjmfhqngnswftgrqvpljsmghszclhhprscrgtmsgbwnbrfhwdsbjrdlhdcqfthwfggpzhjrvmprcjjgrlbmdqczmsltnpnlpldvnsjbdjzbrglpsmbwdhhphcdmjqmssdlgjtmphhjvpjqcdvmdtszldnjtbcffhwrbqpsrsmmsqsbbrnhsrvfhhcmwgcpsfhrfcpvldrwstqtgrcgrqttvdnstlbdblnbtcgdgpbcjdlvwpplsslqdsbpncbwzcmlglbhnrtbwvnbfsgrdlpzszhzmzqtztzgjjjszmjgggdnslqqqrrlwcghpvzwpndbjbhhhpbbtvjvjvjhntqggtmqwdnrtswcchsqpdwnvmfqrgpvbtjpttfhmlgtnsgnpcrhbswfblwrbflqlgtmvzvwprtfntshjcpzshnlnqthvqgpzffsbdbvbhpghwjhqpjwjfbsrzgttsnrpmfrvlwrjrfplrtzpntjbvhdjjfjjzrrvrsrctrtldrsgcgvntgvmtrnqbtmpzvlbwnddbnldbhfcpncqclzqgmbjcnnfdfsmdjdgtmtzbpwqrhrlggrfjhvrrqpjptqdzpwffwnfwmfhfhmcpljmhttbltgltqjhtvbdzgwrppzjwhctnncqthnlrhzgtnffpwnnqqgddjtgtczfmwvbccgvmdfvggdghmssjlflptzrpdncbjcsdlbqjfzmnzdggwdczghfpldsbnvzdjzsbsdnmwptgngmnbqmwbrzqgnqbmfvbzzmsvlswbtvfllqzsjdlvzjmbmlbnljwwjsqzqvfwhzmfdgttfqsdttnjmvspsgvpzwczpvdbvhjjrttghjddrgffqdsmzspnssztjjmwtsrtlrmbqhfbvpzsmnqthssbqrjhvqllrbbpjjllszmnjvzmfsbjtvvvgzrvjwldlbzjdbrsvvgmmjrgsbvtvthjbcqlzqhjqfvbzzvlpmrlmphftwbsvvwnwqfwvsnlvlcdgqjspprmwptcjhbjplzrqlhmvzmbfvwwvgvtqlvgfvzqdprfcjqsnndwqmbhlvfvhdhghtsrhbcrrshlsdtlbwntbcvvjhlcscfcnhbdnngtsccrnfbbqbfggvwnmlfhmqbdqpnzvzcvngcctbrshsdvhstcnwvjgwchzdndgbgjnddrmfgftvvjtdnqhfjmpjcjjtmgztvnwnrjjgmcgcqclnwvclngrtjggrprmfpqcsmdtbswhwrgwgfflqldtpnfdhdltrnhdvqqdsshcrnncdpvbhvlcbjlznjjzsdcjjpqprnzpmvqvrfnhdvhgrsmlcgtbsmrgqqtmblqmqcnmnmbcmmjvnptzbbqmcpwbflhccpfrjvfzmcsttfjtjgrnjbtwbncrsjjgjwdpmvfpqmzqwgjvgvfnpbhgzlzhfsnwmjjjsgcjjcwcfbzrmfzwdmhffqnzscbnfmpnbdbccrsjgqgvftszdbqbstwtwmwfzbtmvdmwvshlcpqnsqfwqlfclmjdjvgszsdblwvnhtwtrjcwsmpgflhdqmndrjvqlhqclmwhscnznmhjtqjlnglpbmgvptjnrpnjlqslssfnbdwlrwrdrtfspsjwzvwzfjpflqgmfnbvzqglrrccsnqtchnhlnrwlwqmqdvwsjmbjbvczvqcgfrrtftcgqsnvzhjlshnlvqmvrjljpsvmdzwqgltmccnplbmplhpvvflhmgfzmhczpwffbzhzhhggzgwtpmlzgrbpfgcnmfdbgcwrbgdpzpwmtwgmjtpmqzzgclmcpjgzdrwqfdwqzzqszrwmbwcqzcjdphqtwlbqmnddsdthgjrlshhmzgrqlhghrbqgzppqbdmmlnthhhmtdcdnbqfvbwdrdgbbpbhtrflwmhwjmzdtwmzhjndbrvbtbqnzbsrrbfdrrdnjltdtsjnpddmqzzbzblzvwctpdhfqgqntczwlgqbvvmhjblztlpptqjscztlrvrglrbsbfdjdfbwjltjcdgmqzgtjqjfgjnnpqgfhsrcggrszhbgfhtjmlgmmvshcpjwmcwcvwhnbjzrzfwtnlsnmdgfpbhvgfrhmhmsmtrnblsmnjclqzfgttrdmvpsvlcvnhmgmltfhffzllfrmdqdpzdmjdmsnmpwbwvrdvwvrcdsfwftfbfvmzjttrqgqlqqdtmcpdpfqfsclcmcdzhjwqrbhmzjjsvpbjzzgfwbjbqmzmtrnflwrqgnbgqnpntqmlngmgcrflgvhpznvrfwzlwrswbdsldfdbhvwvcbwhmpcslhmmbjdzvmgrjwzqqzlrrgddpssqcshzrsvrrzzcbsgvfmrlcgjrcpphctdwtjbmrlnrmrbgsqbjnmnsbjfhszcgdqfflcwvrvnwtvcnhcffvphwmzqpwnwncrbnzrnvjzsrnlzlfdbgztrhvlvpczzqnwsvwvnhgnclhndjsgqdgznjzrtmslmgnjzpj"

func NumberOfCharacterToStartOfPacket(input string, length int) (int, error) {
	runes := []rune(input)
	for i := 0; i <= len(runes)-length; i++ {
		iSOP, err := isStartOfPacket(runes[i:i+length], length)
		if err != nil {
			return 0, err
		}
		if iSOP {
			return i + length, nil
		}
	}
	return 0, fmt.Errorf("no start of packet")
}

func isStartOfPacket(marker []rune, length int) (bool, error) {
	if len(marker) != length {
		return false, fmt.Errorf("marker length is not %d", length)
	}
	r := map[rune]struct{}{}
	for _, m := range marker {
		if _, isDuplicate := r[m]; isDuplicate {
			return false, nil
		}
		r[m] = struct{}{}
	}
	return true, nil
}