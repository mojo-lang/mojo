package org.mojolang.mojo.core;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

import static org.junit.Assert.assertEquals;

@RunWith(JUnit4.class)
public class FieldMaskTreeTest {
	@Test
	public void parentPathReplacesRedundantChildren() {
		FieldMaskTree tree = new FieldMaskTree()
				.addFieldPath("record.name")
				.addFieldPath("record.id")
				.addFieldPath("record");
		assertEquals(FieldMask.newBuilder().addPaths("record").build(), tree.toFieldMask());
	}
}
