package cards

type Card struct {
	Header              *Header      `json:"header,omitempty"`
	Sections            []*Section   `json:"sections,omitempty"`
	SectionDividerStyle DividerStyle `json:"sectionDividerStyle,omitempty"`
	FixedFooter         *FixedFooter `json:"fixedFooter,omitempty"`
}

func NewCard() *Card {
	return &Card{}
}

func (c *Card) SetHeader(h *Header) *Card {
	c.Header = h
	return c
}

func (c *Card) AddSections(s ...*Section) *Card {
	c.Sections = append(c.Sections, s...)
	return c
}

func (c *Card) SetSectionDividerStyle(s DividerStyle) *Card {
	c.SectionDividerStyle = s
	return c
}

func (c *Card) SetFixedFooter(f *FixedFooter) *Card {
	c.FixedFooter = f
	return c
}

type Header struct {
	Title        string    `json:"title,omitempty"`
	Subtitle     string    `json:"subtitle,omitempty"`
	ImageType    ImageType `json:"imageType,omitempty"`
	ImageURL     string    `json:"imageUrl,omitempty"`
	ImageAltText string    `json:"imageAltText,omitempty"`
}

func NewHeader(title string) *Header {
	return &Header{
		Title: title,
	}
}

func (h *Header) SetSubtitle(subtitle string) *Header {
	h.Subtitle = subtitle
	return h
}

func (h *Header) SetImageType(imageType ImageType) *Header {
	h.ImageType = imageType
	return h
}

func (h *Header) SetImageURL(imageURL string) *Header {
	h.ImageURL = imageURL
	return h
}

func (h *Header) SetImageAltText(imageAltText string) *Header {
	h.ImageAltText = imageAltText
	return h
}

type ImageType string

const (
	ImageTypeSquare ImageType = "SQUARE"
	ImageTypeCircle ImageType = "CIRCLE"
)

type Section struct {
	Header                    string    `json:"header,omitempty"`
	Widgets                   []*Widget `json:"widgets,omitempty"`
	Collapsible               bool      `json:"collapsible,omitempty"`
	UncollapsibleWidgetsCount int       `json:"uncollapsibleWidgetsCount,omitempty"`
}

func NewSection() *Section {
	return &Section{}
}

func (s *Section) SetHeader(header string) *Section {
	s.Header = header
	return s
}

func (s *Section) AddWidgets(w ...*Widget) *Section {
	s.Widgets = append(s.Widgets, w...)
	return s
}

func (s *Section) SetCollapsible(collapsible bool) *Section {
	s.Collapsible = collapsible
	return s
}

func (s *Section) SetUncollapsibleWidgetsCount(uncollapsibleWidgetsCount int) *Section {
	s.UncollapsibleWidgetsCount = uncollapsibleWidgetsCount
	return s
}

type Widget struct {
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment,omitempty"`
	TextParagraph       *TextParagraph      `json:"textParagraph,omitempty"`
	Image               *Image              `json:"image,omitempty"`
	DecoratedText       *DecoratedText      `json:"decoratedText,omitempty"`
	ButtonList          *ButtonList         `json:"buttonList,omitempty"`
	TextInput           *TextInput          `json:"textInput,omitempty"`
	SelectionInput      *SelectionInput     `json:"selectionInput,omitempty"`
	DateTimePicker      *DateTimePicker     `json:"dateTimePicker,omitempty"`
	Divider             *Divider            `json:"divider,omitempty"`
	Grid                *Grid               `json:"grid,omitempty"`
	Columns             *Columns            `json:"columns,omitempty"`
}

func NewWidget() *Widget {
	return &Widget{}
}

func (w *Widget) SetHorizontalAlignment(horizontalAlignment HorizontalAlignment) *Widget {
	w.HorizontalAlignment = horizontalAlignment
	return w
}

func (w *Widget) SetTextParagraph(textParagraph *TextParagraph) *Widget {
	w.TextParagraph = textParagraph
	return w
}

func (w *Widget) SetImage(image *Image) *Widget {
	w.Image = image
	return w
}

func (w *Widget) SetDecoratedText(decoratedText *DecoratedText) *Widget {
	w.DecoratedText = decoratedText
	return w
}

func (w *Widget) SetButtonList(buttonList *ButtonList) *Widget {
	w.ButtonList = buttonList
	return w
}

func (w *Widget) SetTextInput(textInput *TextInput) *Widget {
	w.TextInput = textInput
	return w
}

func (w *Widget) SetSelectionInput(selectionInput *SelectionInput) *Widget {
	w.SelectionInput = selectionInput
	return w
}

func (w *Widget) SetDateTimePicker(dateTimePicker *DateTimePicker) *Widget {
	w.DateTimePicker = dateTimePicker
	return w
}

func (w *Widget) SetDivider(divider *Divider) *Widget {
	w.Divider = divider
	return w
}

func (w *Widget) SetGrid(grid *Grid) *Widget {
	w.Grid = grid
	return w
}

func (w *Widget) SetColumns(columns *Columns) *Widget {
	w.Columns = columns
	return w
}

type TextParagraph struct {
	Text string `json:"text,omitempty"`
}

func NewTextParagraph(text string) *TextParagraph {
	return &TextParagraph{
		Text: text,
	}
}

type Image struct {
	ImageURL string   `json:"imageUrl,omitempty"`
	OnClick  *OnClick `json:"onClick,omitempty"`
	AltText  string   `json:"altText,omitempty"`
}

func NewImage(imageURL string) *Image {
	return &Image{
		ImageURL: imageURL,
	}
}

func (i *Image) SetOnClick(onClick *OnClick) *Image {
	i.OnClick = onClick
	return i
}

func (i *Image) SetAltText(altText string) *Image {
	i.AltText = altText
	return i
}

type OnClick struct {
	Action   *Action   `json:"action,omitempty"`
	OpenLink *OpenLink `json:"openLink,omitempty"`
}

func NewOnClick() *OnClick {
	return &OnClick{}
}

func (o *OnClick) SetAction(action *Action) *OnClick {
	o.Action = action
	return o
}

func (o *OnClick) SetOpenLink(openLink *OpenLink) *OnClick {
	o.OpenLink = openLink
	return o
}

type Action struct {
	Function              string             `json:"function,omitempty"`
	Parameters            []*ActionParameter `json:"parameters,omitempty"`
	LoadIndicator         LoadIndicator      `json:"loadIndicator,omitempty"`
	PersistValues         bool               `json:"persistValues,omitempty"`
	Interaction           Interaction        `json:"interaction,omitempty"`
	RequiredWidgets       []string           `json:"requiredWidgets,omitempty"`
	AllWidgetsAreRequired bool               `json:"allWidgetsAreRequired,omitempty"`
}

func NewAction() *Action {
	return &Action{}
}

func (a *Action) SetFunction(function string) *Action {
	a.Function = function
	return a
}

func (a *Action) AddParameters(p ...*ActionParameter) *Action {
	a.Parameters = append(a.Parameters, p...)
	return a
}

func (a *Action) SetLoadIndicator(loadIndicator LoadIndicator) *Action {
	a.LoadIndicator = loadIndicator
	return a
}

func (a *Action) SetPersistValues(persistValues bool) *Action {
	a.PersistValues = persistValues
	return a
}

func (a *Action) SetInteraction(interaction Interaction) *Action {
	a.Interaction = interaction
	return a

}

func (a *Action) AddRequiredWidgets(requiredWidgets ...string) *Action {
	a.RequiredWidgets = append(a.RequiredWidgets, requiredWidgets...)
	return a
}

func (a *Action) SetAllWidgetsAreRequired(allWidgetsAreRequired bool) *Action {
	a.AllWidgetsAreRequired = allWidgetsAreRequired
	return a
}

type ActionParameter struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func NewActionParameter(key, value string) *ActionParameter {
	return &ActionParameter{
		Key:   key,
		Value: value,
	}
}

type LoadIndicator string

const (
	LoadIndicatorSpinner LoadIndicator = "SPINNER"
	LoadIndicatorNone    LoadIndicator = "NONE"
)

type Interaction string

const (
	InteractionUnspecified Interaction = "INTERACTION_UNSPECIFIED"
	InteractionOpenDialog  Interaction = "OPEN_DIALOG"
)

type OpenLink struct {
	URL string `json:"url,omitempty"`
}

func NewOpenLink(url string) *OpenLink {
	return &OpenLink{
		URL: url,
	}
}

type DecoratedText struct {
	StartIcon     *Icon          `json:"startIcon,omitempty"`
	TopLabel      string         `json:"topLabel,omitempty"`
	Text          string         `json:"text,omitempty"`
	WrapText      bool           `json:"wrapText,omitempty"`
	BottomLabel   string         `json:"bottomLabel,omitempty"`
	OnClick       *OnClick       `json:"onClick,omitempty"`
	Button        *Button        `json:"button,omitempty"`
	SwitchControl *SwitchControl `json:"switchControl,omitempty"`
	EndIcon       *Icon          `json:"endIcon,omitempty"`
}

func NewDecoratedText(text string) *DecoratedText {
	return &DecoratedText{
		Text: text,
	}
}

func (d *DecoratedText) SetStartIcon(startIcon *Icon) *DecoratedText {
	d.StartIcon = startIcon
	return d
}

func (d *DecoratedText) SetTopLabel(topLabel string) *DecoratedText {
	d.TopLabel = topLabel
	return d
}

func (d *DecoratedText) SetWrapText(wrapText bool) *DecoratedText {
	d.WrapText = wrapText
	return d
}

func (d *DecoratedText) SetBottomLabel(bottomLabel string) *DecoratedText {
	d.BottomLabel = bottomLabel
	return d
}

func (d *DecoratedText) SetOnClick(onClick *OnClick) *DecoratedText {
	d.OnClick = onClick
	return d
}

func (d *DecoratedText) SetButton(button *Button) *DecoratedText {
	d.Button = button
	return d
}

func (d *DecoratedText) SetSwitchControl(switchControl *SwitchControl) *DecoratedText {
	d.SwitchControl = switchControl
	return d
}

func (d *DecoratedText) SetEndIcon(endIcon *Icon) *DecoratedText {
	d.EndIcon = endIcon
	return d
}

type Icon struct {
	AltText      string        `json:"altText,omitempty"`
	ImageType    ImageType     `json:"imageType,omitempty"`
	KnownIcon    string        `json:"knownIcon,omitempty"`
	IconURL      string        `json:"iconUrl,omitempty"`
	MaterialIcon *MaterialIcon `json:"materialIcon,omitempty"`
}

func NewIcon() *Icon {
	return &Icon{}
}

func (i *Icon) SetAltText(altText string) *Icon {
	i.AltText = altText
	return i
}

func (i *Icon) SetImageType(imageType ImageType) *Icon {
	i.ImageType = imageType
	return i
}

func (i *Icon) SetKnownIcon(knownIcon string) *Icon {
	i.KnownIcon = knownIcon
	return i
}

func (i *Icon) SetIconURL(iconURL string) *Icon {
	i.IconURL = iconURL
	return i
}

func (i *Icon) SetMaterialIcon(materialIcon *MaterialIcon) *Icon {
	i.MaterialIcon = materialIcon
	return i
}

type MaterialIcon struct {
	Name   string `json:"name,omitempty"`
	Fill   bool   `json:"fill,omitempty"`
	Weight int    `json:"weight,omitempty"`
	Grade  int    `json:"grade,omitempty"`
}

func NewMaterialIcon(name string) *MaterialIcon {
	return &MaterialIcon{
		Name: name,
	}
}

func (m *MaterialIcon) SetFill(fill bool) *MaterialIcon {
	m.Fill = fill
	return m
}

func (m *MaterialIcon) SetWeight(weight int) *MaterialIcon {
	m.Weight = weight
	return m
}

func (m *MaterialIcon) SetGrade(grade int) *MaterialIcon {
	m.Grade = grade
	return m
}

type Button struct {
	Text     string   `json:"text,omitempty"`
	Icon     *Icon    `json:"icon,omitempty"`
	Color    *Color   `json:"color,omitempty"`
	OnClick  *OnClick `json:"onClick,omitempty"`
	Disabled bool     `json:"disabled,omitempty"`
	AltText  string   `json:"altText,omitempty"`
}

func NewButton(text string) *Button {
	return &Button{
		Text: text,
	}
}

func (b *Button) SetIcon(icon *Icon) *Button {
	b.Icon = icon
	return b
}

func (b *Button) SetColor(color *Color) *Button {
	b.Color = color
	return b
}

func (b *Button) SetOnClick(onClick *OnClick) *Button {
	b.OnClick = onClick
	return b
}

func (b *Button) SetDisabled(disabled bool) *Button {
	b.Disabled = disabled
	return b
}

func (b *Button) SetAltText(altText string) *Button {
	b.AltText = altText
	return b
}

type Color struct {
	Red   float64 `json:"red,omitempty"`
	Green float64 `json:"green,omitempty"`
	Blue  float64 `json:"blue,omitempty"`
	Alpha float64 `json:"alpha,omitempty"`
}

func NewColor(red, green, blue, alpha float64) *Color {
	return &Color{
		Red:   red,
		Green: green,
		Blue:  blue,
		Alpha: alpha,
	}
}

type SwitchControl struct {
	Name           string      `json:"name,omitempty"`
	Value          string      `json:"value,omitempty"`
	Selected       bool        `json:"selected,omitempty"`
	OnChangeAction *Action     `json:"onChangeAction,omitempty"`
	ControlType    ControlType `json:"controlType,omitempty"`
}

func NewSwitchControl() *SwitchControl {
	return &SwitchControl{}
}

func (s *SwitchControl) SetName(name string) *SwitchControl {
	s.Name = name
	return s
}

func (s *SwitchControl) SetValue(value string) *SwitchControl {
	s.Value = value
	return s
}

func (s *SwitchControl) SetSelected(selected bool) *SwitchControl {
	s.Selected = selected
	return s
}

func (s *SwitchControl) SetOnChangeAction(onChangeAction *Action) *SwitchControl {
	s.OnChangeAction = onChangeAction
	return s
}

func (s *SwitchControl) SetControlType(controlType ControlType) *SwitchControl {
	s.ControlType = controlType
	return s
}

type ControlType string

const (
	ControlTypeSwitch   ControlType = "SWITCH"
	ControlTypeCheckBox ControlType = "CHECK_BOX"
)

type ButtonList struct {
	Buttons []*Button `json:"buttons,omitempty"`
}

func NewButtonList() *ButtonList {
	return &ButtonList{}
}

func (b *ButtonList) AddButtons(buttons ...*Button) *ButtonList {
	b.Buttons = append(b.Buttons, buttons...)
	return b
}

type TextInput struct {
	Name               string       `json:"name,omitempty"`
	Label              string       `json:"label,omitempty"`
	HintText           string       `json:"hintText,omitempty"`
	Value              string       `json:"value,omitempty"`
	Type               Type         `json:"type,omitempty"`
	OnChangeAction     *Action      `json:"onChangeAction,omitempty"`
	InitialSuggestions *Suggestions `json:"initialSuggestions,omitempty"`
	Validation         *Validation  `json:"validation,omitempty"`
	PlaceholderText    string       `json:"placeholderText,omitempty"`
}

func NewTextInput() *TextInput {
	return &TextInput{}
}

func (t *TextInput) SetName(name string) *TextInput {
	t.Name = name
	return t
}

func (t *TextInput) SetLabel(label string) *TextInput {
	t.Label = label
	return t
}

func (t *TextInput) SetHintText(hintText string) *TextInput {
	t.HintText = hintText
	return t
}

func (t *TextInput) SetValue(value string) *TextInput {
	t.Value = value
	return t
}

func (t *TextInput) SetType(inputType Type) *TextInput {
	t.Type = inputType
	return t
}

func (t *TextInput) SetOnChangeAction(onChangeAction *Action) *TextInput {
	t.OnChangeAction = onChangeAction
	return t
}

func (t *TextInput) SetInitialSuggestions(initialSuggestions *Suggestions) *TextInput {
	t.InitialSuggestions = initialSuggestions
	return t
}

func (t *TextInput) SetValidation(validation *Validation) *TextInput {
	t.Validation = validation
	return t
}

func (t *TextInput) SetPlaceholderText(placeholderText string) *TextInput {
	t.PlaceholderText = placeholderText
	return t
}

type Type string

const (
	TypeSingleLine   Type = "SINGLE_LINE"
	TypeMultipleLine Type = "MULTIPLE_LINE"
)

type Suggestions struct {
	Items []*SuggestionItem `json:"items,omitempty"`
}

func NewSuggestions(items ...*SuggestionItem) *Suggestions {
	return &Suggestions{
		Items: items,
	}
}

type SuggestionItem struct {
	Text string `json:"text,omitempty"`
}

func NewSuggestionItem(text string) *SuggestionItem {
	return &SuggestionItem{
		Text: text,
	}
}

type Validation struct {
	CharacterLimit int       `json:"characterLimit,omitempty"`
	InputType      InputType `json:"enum,omitempty"`
}

func NewValidation() *Validation {
	return &Validation{}
}

func (v *Validation) SetCharacterLimit(characterLimit int) *Validation {
	v.CharacterLimit = characterLimit
	return v
}

func (v *Validation) SetInputType(inputType InputType) *Validation {
	v.InputType = inputType
	return v
}

type InputType string

const (
	InputTypeText        InputType = "TEXT"
	InputTypeInteger     InputType = "INTEGER"
	InputTypeFloat       InputType = "FLOAT"
	InputTypeEmail       InputType = "EMAIL"
	InputTypeEmojiPicker InputType = "EMOJI_PICKER"
)

type SelectionInput struct {
	Name                        string              `json:"name,omitempty"`
	Label                       string              `json:"label,omitempty"`
	Type                        SelectionType       `json:"type,omitempty"`
	Items                       []*SelectionItem    `json:"items,omitempty"`
	OnChangeAction              *Action             `json:"onChangeAction,omitempty"`
	MultiSelectMaxSelectedItems int                 `json:"multiSelectMaxSelectedItems,omitempty"`
	MultiSelectMinQueryLength   int                 `json:"multiSelectMinQueryLength,omitempty"`
	Validation                  *Validation         `json:"validation,omitempty"`
	ExternalDataSource          *Action             `json:"externalDataSource,omitempty"`
	PlatformDataSource          *PlatformDataSource `json:"platformDataSource,omitempty"`
}

func NewSelectionInput() *SelectionInput {
	return &SelectionInput{}
}

func (s *SelectionInput) SetName(name string) *SelectionInput {
	s.Name = name
	return s
}

func (s *SelectionInput) SetLabel(label string) *SelectionInput {
	s.Label = label
	return s
}

func (s *SelectionInput) SetType(selectionType SelectionType) *SelectionInput {
	s.Type = selectionType
	return s
}

func (s *SelectionInput) AddItems(items ...*SelectionItem) *SelectionInput {
	s.Items = append(s.Items, items...)
	return s
}

func (s *SelectionInput) SetOnChangeAction(onChangeAction *Action) *SelectionInput {
	s.OnChangeAction = onChangeAction
	return s
}

func (s *SelectionInput) SetMultiSelectMaxSelectedItems(multiSelectMaxSelectedItems int) *SelectionInput {
	s.MultiSelectMaxSelectedItems = multiSelectMaxSelectedItems
	return s
}

func (s *SelectionInput) SetMultiSelectMinQueryLength(multiSelectMinQueryLength int) *SelectionInput {
	s.MultiSelectMinQueryLength = multiSelectMinQueryLength
	return s
}

func (s *SelectionInput) SetValidation(validation *Validation) *SelectionInput {
	s.Validation = validation
	return s
}

type SelectionType string

const (
	SelectionTypeCheckBox    SelectionType = "CHECK_BOX"
	SelectionTypeRadio       SelectionType = "RADIO_BUTTON"
	SelectionTypeDropdown    SelectionType = "DROPDOWN"
	SelectionTypeSwitch      SelectionType = "SWITCH"
	SelectionTypeMultiSelect SelectionType = "MULTI_SELECT"
)

type SelectionItem struct {
	Text         string `json:"text,omitempty"`
	Value        string `json:"value,omitempty"`
	Selected     bool   `json:"selected,omitempty"`
	StartIconURI string `json:"startIconUri,omitempty"`
	BottomText   string `json:"bottomText,omitempty"`
}

func NewSelectionItem() *SelectionItem {
	return &SelectionItem{}
}

func (s *SelectionItem) SetText(text string) *SelectionItem {
	s.Text = text
	return s
}

func (s *SelectionItem) SetValue(value string) *SelectionItem {
	s.Value = value
	return s
}

func (s *SelectionItem) SetSelected(selected bool) *SelectionItem {
	s.Selected = selected
	return s
}

func (s *SelectionItem) SetStartIconURI(startIconURI string) *SelectionItem {
	s.StartIconURI = startIconURI
	return s
}

func (s *SelectionItem) SetBottomText(bottomText string) *SelectionItem {
	s.BottomText = bottomText
	return s
}

type PlatformDataSource struct {
	CommonDataSource  CommonDataSource         `json:"commonDataSource,omitempty"`
	HostAppDataSource *HostAppDataSourceMarkup `json:"hostAppDataSource,omitempty"`
}

func NewPlatformDataSource() *PlatformDataSource {
	return &PlatformDataSource{}
}

func (p *PlatformDataSource) SetCommonDataSource(commonDataSource CommonDataSource) *PlatformDataSource {
	p.CommonDataSource = commonDataSource
	return p
}

func (p *PlatformDataSource) SetHostAppDataSource(hostAppDataSource *HostAppDataSourceMarkup) *PlatformDataSource {
	p.HostAppDataSource = hostAppDataSource
	return p
}

type CommonDataSource string

const (
	CommonDataSourceUser CommonDataSource = "USER"
)

type HostAppDataSourceMarkup struct {
	ChatDataSource *ChatClientDataSourceMarkup `json:"chatDataSource,omitempty"`
}

func NewHostAppDataSourceMarkup() *HostAppDataSourceMarkup {
	return &HostAppDataSourceMarkup{}
}

func (h *HostAppDataSourceMarkup) SetChatDataSource(chatDataSource *ChatClientDataSourceMarkup) *HostAppDataSourceMarkup {
	h.ChatDataSource = chatDataSource
	return h
}

type ChatClientDataSourceMarkup struct {
	SpaceDataSource *SpaceDataSource `json:"spaceDataSource,omitempty"`
}

func NewChatClientDataSourceMarkup() *ChatClientDataSourceMarkup {
	return &ChatClientDataSourceMarkup{}
}

func (c *ChatClientDataSourceMarkup) SetSpaceDataSource(spaceDataSource *SpaceDataSource) *ChatClientDataSourceMarkup {
	c.SpaceDataSource = spaceDataSource
	return c
}

type SpaceDataSource struct {
	DefaultToCurrentSpace bool `json:"defaultToCurrentSpace,omitempty"`
}

func NewSpaceDataSource() *SpaceDataSource {
	return &SpaceDataSource{}
}

func (s *SpaceDataSource) SetDefaultToCurrentSpace(defaultToCurrentSpace bool) *SpaceDataSource {
	s.DefaultToCurrentSpace = defaultToCurrentSpace
	return s
}

type DateTimePicker struct {
	Name               string             `json:"name,omitempty"`
	Label              string             `json:"label,omitempty"`
	Type               DateTimePickerType `json:"type,omitempty"`
	ValueMsEpoch       string             `json:"valueMsEpoch,omitempty"`
	TimezoneOffsetDate int                `json:"timezoneOffsetDate,omitempty"`
	OnChangeAction     *Action            `json:"onChangeAction,omitempty"`
	Validation         *Validation        `json:"validation,omitempty"`
}

func NewDateTimePicker() *DateTimePicker {
	return &DateTimePicker{}
}

func (d *DateTimePicker) SetName(name string) *DateTimePicker {
	d.Name = name
	return d
}

func (d *DateTimePicker) SetLabel(label string) *DateTimePicker {
	d.Label = label
	return d
}

func (d *DateTimePicker) SetType(dateTimePickerType DateTimePickerType) *DateTimePicker {
	d.Type = dateTimePickerType
	return d
}

func (d *DateTimePicker) SetValueMsEpoch(valueMsEpoch string) *DateTimePicker {
	d.ValueMsEpoch = valueMsEpoch
	return d
}

func (d *DateTimePicker) SetTimezoneOffsetDate(timezoneOffsetDate int) *DateTimePicker {
	d.TimezoneOffsetDate = timezoneOffsetDate
	return d
}

func (d *DateTimePicker) SetOnChangeAction(onChangeAction *Action) *DateTimePicker {
	d.OnChangeAction = onChangeAction
	return d
}

func (d *DateTimePicker) SetValidation(validation *Validation) *DateTimePicker {
	d.Validation = validation
	return d
}

type DateTimePickerType string

const (
	DateTimePickerDateAndTime DateTimePickerType = "DATE_AND_TIME"
	DateTimePickerDateOnly    DateTimePickerType = "DATE_ONLY"
	DateTimePickerTimeOnly    DateTimePickerType = "TIME_ONLY"
)

type Divider struct {
}

type Grid struct {
	Title       string       `json:"title,omitempty"`
	Items       []*GridItem  `json:"items,omitempty"`
	BorderStyle *BorderStyle `json:"borderStyle,omitempty"`
	ColumnCount int          `json:"columnCount,omitempty"`
	OnClick     *OnClick     `json:"onClick,omitempty"`
}

func NewGrid() *Grid {
	return &Grid{}
}

func (g *Grid) SetTitle(title string) *Grid {
	g.Title = title
	return g
}

func (g *Grid) AddItems(items ...*GridItem) *Grid {
	g.Items = append(g.Items, items...)
	return g
}

func (g *Grid) SetBorderStyle(borderStyle *BorderStyle) *Grid {
	g.BorderStyle = borderStyle
	return g
}

func (g *Grid) SetColumnCount(columnCount int) *Grid {
	g.ColumnCount = columnCount
	return g
}

func (g *Grid) SetOnClick(onClick *OnClick) *Grid {
	g.OnClick = onClick
	return g
}

type GridItem struct {
	ID       string          `json:"id,omitempty"`
	Image    *ImageComponent `json:"image,omitempty"`
	Title    string          `json:"title,omitempty"`
	Subtitle string          `json:"subtitle,omitempty"`
	Layout   GridItemLayout  `json:"layout,omitempty"`
}

func NewGridItem() *GridItem {
	return &GridItem{}
}

func (g *GridItem) SetID(id string) *GridItem {
	g.ID = id
	return g
}

func (g *GridItem) SetImage(image *ImageComponent) *GridItem {
	g.Image = image
	return g
}

func (g *GridItem) SetTitle(title string) *GridItem {
	g.Title = title
	return g
}

func (g *GridItem) SetSubtitle(subtitle string) *GridItem {
	g.Subtitle = subtitle
	return g
}

func (g *GridItem) SetLayout(layout GridItemLayout) *GridItem {
	g.Layout = layout
	return g
}

type ImageComponent struct {
	ImageURI    string          `json:"imageUri,omitempty"`
	AltText     string          `json:"altText,omitempty"`
	CropStyle   *ImageCropStyle `json:"cropStyle,omitempty"`
	BorderStyle *BorderStyle    `json:"borderStyle,omitempty"`
}

func NewImageComponent() *ImageComponent {
	return &ImageComponent{}
}

func (i *ImageComponent) SetImageURI(imageURI string) *ImageComponent {
	i.ImageURI = imageURI
	return i
}

func (i *ImageComponent) SetAltText(altText string) *ImageComponent {
	i.AltText = altText
	return i
}

func (i *ImageComponent) SetCropStyle(cropStyle *ImageCropStyle) *ImageComponent {
	i.CropStyle = cropStyle
	return i
}

func (i *ImageComponent) SetBorderStyle(borderStyle *BorderStyle) *ImageComponent {
	i.BorderStyle = borderStyle
	return i
}

type ImageCropStyle struct {
	Type        ImageCropType `json:"type,omitempty"`
	AspectRatio float64       `json:"aspectRatio,omitempty"`
}

func NewImageCropStyle() *ImageCropStyle {
	return &ImageCropStyle{}
}

func (i *ImageCropStyle) SetType(imageCropType ImageCropType) *ImageCropStyle {
	i.Type = imageCropType
	return i
}

func (i *ImageCropStyle) SetAspectRatio(aspectRatio float64) *ImageCropStyle {
	i.AspectRatio = aspectRatio
	return i
}

type ImageCropType string

const (
	ImageCropSquare          ImageCropType = "SQUARE"
	ImageCropCircle          ImageCropType = "CIRCLE"
	ImageCropRectangleCustom ImageCropType = "RECTANGLE_CUSTOM"
	ImageCropRectangle4_3    ImageCropType = "RECTANGLE_4_3"
)

type BorderStyle struct {
	Type         BorderType `json:"type,omitempty"`
	StrokeColor  *Color     `json:"strokeColor,omitempty"`
	CornerRadius int        `json:"cornerRadius,omitempty"`
}

func NewBorderStyle() *BorderStyle {
	return &BorderStyle{}
}

func (b *BorderStyle) SetType(borderType BorderType) *BorderStyle {
	b.Type = borderType
	return b
}

func (b *BorderStyle) SetStrokeColor(strokeColor *Color) *BorderStyle {
	b.StrokeColor = strokeColor
	return b
}

func (b *BorderStyle) SetCornerRadius(cornerRadius int) *BorderStyle {
	b.CornerRadius = cornerRadius
	return b
}

type BorderType string

const (
	BorderNoBorder BorderType = "NO_BORDER"
	BorderStroke   BorderType = "STROKE"
)

type GridItemLayout string

const (
	GridItemTextBelow ImageCropType = "TEXT_BELOW"
	GridItemTextAbove ImageCropType = "TEXT_ABOVE"
)

type Columns struct {
	ColumnItems []*Column `json:"columnItems,omitempty"`
}

func NewColumns(c ...*Column) *Columns {
	return &Columns{
		ColumnItems: c,
	}
}

func (c *Columns) AddColumnItems(columnItems ...*Column) *Columns {
	c.ColumnItems = append(c.ColumnItems, columnItems...)
	return c
}

type Column struct {
	HorizontalSizeStyle HorizontalSizeStyle `json:"horizontalSizeStyle,omitempty"`
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment,omitempty"`
	VerticalAlignment   VerticalAlignment   `json:"verticalAlignment,omitempty"`
	Widgets             []*Widgets          `json:"widgets,omitempty"`
}

func NewColumn() *Column {
	return &Column{}
}

func (c *Column) SetHorizontalSizeStyle(horizontalSizeStyle HorizontalSizeStyle) *Column {
	c.HorizontalSizeStyle = horizontalSizeStyle
	return c
}

func (c *Column) SetHorizontalAlignment(horizontalAlignment HorizontalAlignment) *Column {
	c.HorizontalAlignment = horizontalAlignment
	return c
}

func (c *Column) SetVerticalAlignment(verticalAlignment VerticalAlignment) *Column {
	c.VerticalAlignment = verticalAlignment
	return c
}

func (c *Column) AddWidgets(w ...*Widgets) *Column {
	c.Widgets = append(c.Widgets, w...)
	return c
}

type HorizontalSizeStyle string

const (
	HorizontalSizeFillAvailableSpace HorizontalSizeStyle = "FILL_AVAILABLE_SPACE"
	HorizontalSizeFillMinimumSpace   HorizontalSizeStyle = "FILL_MINIMUM_SPACE"
)

type HorizontalAlignment string

const (
	HorizontalAlignmentStart  HorizontalAlignment = "START"
	HorizontalAlignmentCenter HorizontalAlignment = "CENTER"
	HorizontalAlignmentEnd    HorizontalAlignment = "END"
)

type VerticalAlignment string

const (
	VerticalAlignmentCenter VerticalAlignment = "CENTER"
	VerticalAlignmentTop    VerticalAlignment = "TOP"
	VerticalAlignmentBottom VerticalAlignment = "BOTTOM"
)

type Widgets struct {
	TextParagraph  *TextParagraph  `json:"textParagraph,omitempty"`
	Image          *Image          `json:"image,omitempty"`
	DecoratedText  *DecoratedText  `json:"decoratedText,omitempty"`
	ButtonList     *ButtonList     `json:"buttonList,omitempty"`
	TextInput      *TextInput      `json:"textInput,omitempty"`
	SelectionInput *SelectionInput `json:"selectionInput,omitempty"`
	DateTimePicker *DateTimePicker `json:"dateTimePicker,omitempty"`
}

func NewWidgets() *Widgets {
	return &Widgets{}
}

func (w *Widgets) SetTextParagraph(textParagraph *TextParagraph) *Widgets {
	w.TextParagraph = textParagraph
	return w
}

func (w *Widgets) SetImage(image *Image) *Widgets {
	w.Image = image
	return w
}

func (w *Widgets) SetDecoratedText(decoratedText *DecoratedText) *Widgets {
	w.DecoratedText = decoratedText
	return w
}

func (w *Widgets) SetButtonList(buttonList *ButtonList) *Widgets {
	w.ButtonList = buttonList
	return w
}

func (w *Widgets) SetTextInput(textInput *TextInput) *Widgets {
	w.TextInput = textInput
	return w
}

func (w *Widgets) SetSelectionInput(selectionInput *SelectionInput) *Widgets {
	w.SelectionInput = selectionInput
	return w
}

func (w *Widgets) SetDateTimePicker(dateTimePicker *DateTimePicker) *Widgets {
	w.DateTimePicker = dateTimePicker
	return w
}

type DividerStyle string

const (
	DividerSolid DividerStyle = "SOLID_DIVIDER"
	DividerNone  DividerStyle = "NO_DIVIDER"
)

type FixedFooter struct {
	PrimaryButton   *Button `json:"primaryButton,omitempty"`
	SecondaryButton *Button `json:"secondaryButton,omitempty"`
}

func NewFixedFooter() *FixedFooter {
	return &FixedFooter{}
}

func (f *FixedFooter) SetPrimaryButton(primaryButton *Button) *FixedFooter {
	f.PrimaryButton = primaryButton
	return f
}

func (f *FixedFooter) SetSecondaryButton(secondaryButton *Button) *FixedFooter {
	f.SecondaryButton = secondaryButton
	return f
}
