package internal

import (
	"bytes"
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalChildrenDescandantsImpl_Children(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx           context.Context
		contentID     string
		expand        []string
		parentVersion int
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:           context.TODO(),
				contentID:     "100100101",
				expand:        []string{"attachment", "comments"},
				parentVersion: 12,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/child?expand=attachment%2Ccomments&parentVersion=12",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentChildrenScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:           context.TODO(),
				contentID:     "100100101",
				expand:        []string{"attachment", "comments"},
				parentVersion: 12,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/child?expand=attachment%2Ccomments&parentVersion=12",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Children(testCase.args.ctx, testCase.args.contentID, testCase.args.expand,
				testCase.args.parentVersion)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalChildrenDescandantsImpl_ChildrenByType(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                    context.Context
		contentID, contentType string
		expand                 []string
		parentVersion          int
		startAt, maxResults    int
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:           context.TODO(),
				contentID:     "100100101",
				contentType:   "blogpost",
				expand:        []string{"attachment", "comments"},
				parentVersion: 12,
				startAt:       50,
				maxResults:    25,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/child/blogpost?expand=attachment%2Ccomments&limit=25&parentVersion=12&start=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:           context.TODO(),
				contentID:     "100100101",
				contentType:   "blogpost",
				expand:        []string{"attachment", "comments"},
				parentVersion: 12,
				startAt:       50,
				maxResults:    25,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/child/blogpost?expand=attachment%2Ccomments&limit=25&parentVersion=12&start=50",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},

		{
			name: "when the content type is not provided",
			args: args{
				ctx:       context.TODO(),
				contentID: "11929292",
			},
			wantErr: true,
			Err:     model.ErrNoContentTypeError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.ChildrenByType(testCase.args.ctx, testCase.args.contentID,
				testCase.args.contentType, testCase.args.parentVersion, testCase.args.expand, testCase.args.startAt,
				testCase.args.maxResults)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalChildrenDescandantsImpl_Descendants(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx       context.Context
		contentID string
		expand    []string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:       context.TODO(),
				contentID: "100100101",
				expand:    []string{"attachment", "comments"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/descendant?expand=attachment%2Ccomments",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentChildrenScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.TODO(),
				contentID: "100100101",
				expand:    []string{"attachment", "comments"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/descendant?expand=attachment%2Ccomments",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Descendants(testCase.args.ctx, testCase.args.contentID, testCase.args.expand)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})
	}
}

func Test_internalChildrenDescandantsImpl_DescendantsByType(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                           context.Context
		contentID, contentType, depth string
		expand                        []string
		startAt, maxResults           int
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:         context.TODO(),
				contentID:   "100100101",
				contentType: "blogpost",
				expand:      []string{"attachment", "comments"},
				startAt:     50,
				maxResults:  25,
				depth:       "root",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/descendant/blogpost?depth=root&expand=attachment%2Ccomments&limit=25&start=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				contentID:   "100100101",
				contentType: "blogpost",
				expand:      []string{"attachment", "comments"},
				startAt:     50,
				maxResults:  25,
				depth:       "root",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/descendant/blogpost?depth=root&expand=attachment%2Ccomments&limit=25&start=50",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},

		{
			name: "when the content type is not provided",
			args: args{
				ctx:       context.TODO(),
				contentID: "11929292",
			},
			wantErr: true,
			Err:     model.ErrNoContentTypeError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.DescendantsByType(testCase.args.ctx, testCase.args.contentID,
				testCase.args.contentType, testCase.args.depth, testCase.args.expand, testCase.args.startAt,
				testCase.args.maxResults)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalChildrenDescandantsImpl_CopyHierarchy(t *testing.T) {

	payloadMocked := &model.CopyOptionsScheme{
		CopyAttachments:    true,
		CopyPermissions:    true,
		CopyProperties:     true,
		CopyLabels:         true,
		CopyCustomContents: true,
		DestinationPageID:  "223322",
		TitleOptions: &model.CopyTitleOptionScheme{
			Prefix:  "copy-",
			Replace: "test",
		},
		PageTitle: "Test Title",
	}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx       context.Context
		contentID string
		options   *model.CopyOptionsScheme
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:       context.TODO(),
				contentID: "100100101",
				options:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/100100101/pagehierarchy/copy",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.TaskScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.TODO(),
				contentID: "100100101",
				options:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/100100101/pagehierarchy/copy",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.CopyHierarchy(testCase.args.ctx, testCase.args.contentID,
				testCase.args.options)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalChildrenDescandantsImpl_CopyPage(t *testing.T) {

	payloadMocked := &model.CopyOptionsScheme{
		CopyAttachments:    true,
		CopyPermissions:    true,
		CopyProperties:     true,
		CopyLabels:         true,
		CopyCustomContents: true,
		DestinationPageID:  "223322",
		TitleOptions: &model.CopyTitleOptionScheme{
			Prefix:  "copy-",
			Replace: "test",
		},
		PageTitle: "Test Title",
	}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx       context.Context
		contentID string
		expand    []string
		options   *model.CopyOptionsScheme
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:       context.TODO(),
				contentID: "100100101",
				options:   payloadMocked,
				expand:    []string{"childTypes.all"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/100100101/copy?expand=childTypes.all",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.TODO(),
				contentID: "100100101",
				options:   payloadMocked,
				expand:    []string{"childTypes.all"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/100100101/copy?expand=childTypes.all",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.CopyPage(testCase.args.ctx, testCase.args.contentID,
				testCase.args.expand, testCase.args.options)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}