package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Equal(t, []string{}, Top10(""))
	})

	t.Run("no words in whitespace-only string", func(t *testing.T) {
		require.Equal(t, []string{}, Top10(" \t\n "))
	})

	t.Run("splits words by mixed whitespace separators", func(t *testing.T) {
		expected := []string{"one", "three", "two"}

		require.Equal(t, expected, Top10("one\tone\n two  three"))
	})

	t.Run("keeps README example as exact tokens", func(t *testing.T) {
		expected := []string{"and", "one", "cat", "cats", "dog", "dog,two", "man"}

		require.Equal(t, expected, Top10("cat and dog, one dog,two cats and one man"))
	})

	t.Run("normalizes case and edge punctuation", func(t *testing.T) {
		expected := []string{"нога"}

		require.Equal(t, expected, Top10("Нога нога нога, 'НОГА'"))
	})

	t.Run("keeps inner punctuation intact", func(t *testing.T) {
		expected := []string{"dog,cat", "dog...cat", "dogcat"}

		require.Equal(t, expected, Top10("dog,cat dog...cat dogcat"))
	})

	t.Run("drops single hyphen but keeps longer hyphen-only tokens", func(t *testing.T) {
		expected := []string{"--", "-------"}

		require.Equal(t, expected, Top10("- -- ------- -"))
	})

	t.Run("trims edge punctuation before applying hyphen-only rule", func(t *testing.T) {
		expected := []string{"--", "---"}

		require.Equal(t, expected, Top10("'-' '(--)' '---'"))
	})

	t.Run("keeps lexicographically first ten after normalization tie", func(t *testing.T) {
		expected := []string{"w01", "w02", "w03", "w04", "w05", "w06", "w07", "w08", "w09", "w10"}

		require.Equal(t, expected, Top10("w12! w11? w10. w09, w08: w07; w06 w05 w04 w03 w02 w01"))
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}
