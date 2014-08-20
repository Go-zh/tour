/* Copyright 2012 The Go Authors.   All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
'use strict';

angular.module('tour.values', []).

// List of modules with description and lessons in it.
value('tableOfContents', [{
    'id': 'mechanics',
    'title': '指南的使用',
    'description': '<p>欢迎来到 <a href="http://golang.org">Go 编程语言</a>指南。本指南涵盖了该语言的大部分重要特性，主要包括：</p>',
    'lessons': ['welcome']
}, {
    'id': 'basics',
    'title': '基础',
    'description': '<p>一开始，将学习关于语言的所有基础内容。</p><p>定义变量、调用函数、以及在你学习下一课之前所需要了解的全部内容。</p>',
    'lessons': ['basics', 'flowcontrol', 'moretypes']
}, {
    'id': 'methods',
    'title': '方法和接口',
    'description': '<p>学习如何为类型定义方法；如何定义接口；以及如何将所有内容贯通起来。</p>',
    'lessons': ['methods']
}, {
    'id': 'concurrency',
    'title': '并发',
    'description': '<p>作为语言的核心部分，Go 提供了并发的特性。</p><p>这一部分概览了 goroutein 和 channel，以及如何使用它们来实现不同的并发模式。</p>',
    'lessons': ['concurrency']
}]).

// translation
value('translation', {
    'off': 'off',
    'on': 'on',
    'syntax': '语法高亮',
    'lineno': '行号',
    'reset': '重置',
    'format': '格式化源代码',
    'kill': '杀死进程',
    'run': '运行',
    'compile': '编译并运行',
    'more': '选项',
    'toc': '目录',
    'prev': '向前',
    'next': '向后',
    'waiting': '等待远端服务器响应...',
    'errcomm': '与远端服务器通讯失败。',
}).

// Config for codemirror plugin
value('ui.config', {
    codemirror: {
        mode: 'text/x-go',
        matchBrackets: true,
        lineNumbers: true,
        autofocus: true,
        indentWithTabs: true,
        lineWrapping: true,
        extraKeys: {
            'Shift-Enter': function() {
                $('#run').click();
            },
            'Ctrl-Enter': function() {
                $('#format').click();
            },
            'PageDown': function() {
                return false;
            },
            'PageUp': function() {
                return false;
            },
        }
    }
});
